package api

import (
	"net/http"
	"reflect"

	"github.com/kolo/xmlrpc"
	"github.com/xoes/go-odoo/types"
)


type Config struct {
	HostURL   		string
	AdminPassword   string
	Transport 		http.RoundTripper
	Session
}

type Session struct {
	DbName   string
	Username string
	Password string
	UserID   int
}

type CreateDatabaseConfig struct {
	dbName string
	demo bool
	lang string
	userPassword string
	login string
	countryCode string
}

func (c *Config) NewClient() (*xmlrpc.Client, error) {
	client, err := xmlrpc.NewClient(c.HostURL + "/xmlrpc/2/object", c.Transport)
	if err != nil {
		return nil, err
	}
	return client, err
}

func (c *Config) Login(s Session) error {
	var uid int
	endpointURL := c.HostURL + "/xmlrpc/2/common"
	client, err := xmlrpc.NewClient(endpointURL, c.Transport)
	if err != nil {
		return err
	}
	err = client.Call("authenticate", []interface{}{s.DbName, s.Username, s.Password, ""}, &uid)
	if err != nil {
		return err
	}
	s.UserID = uid
	c.Session = s
	return err
}

func (c *Config) CreateDatabase(dbConfig *CreateDatabaseConfig) error {
	var reply bool
	endpointURL := c.HostURL + "/xmlrpc/2/db"
	client, err := xmlrpc.NewClient(endpointURL, c.Transport)
	if err != nil {
		return err
	}
	err = client.Call("create_database", []interface{}{dbConfig}, &reply)
	if err != nil {
		return err
	}
	return err
}

func (c *Config) DropDatabase(dbName string) error {
	var reply bool
	endpointURL := c.HostURL + "/xmlrpc/2/db"
	client, err := xmlrpc.NewClient(endpointURL, c.Transport)
	if err != nil {
		return err
	}
	err = client.Call("drop", []interface{}{c.AdminPassword, dbName}, &reply)
	if err != nil {
		return err
	}
	return err
}

// Low-level functions
func (c *Config) Create(model string, args []interface{}, elem interface{}) error {
	return c.DoRequest("create", model, args, nil, elem)
}

func (c *Config) Update(model string, args []interface{}) error {
	return c.DoRequest("write", model, args, nil, nil)
}

func (c *Config) Delete(model string, args []interface{}) error {
	return c.DoRequest("unlink", model, args, nil, nil)
}

func (c *Config) Search(model string, args []interface{}, options interface{}, elem interface{}) error {
	return c.DoRequest("search", model, args, options, elem)
}

func (c *Config) Read(model string, args []interface{}, options interface{}, elem interface{}) error {
	ne := elem.(types.Type).NilableType_()
	err := c.DoRequest("read", model, args, options, ne)
	if err == nil {
		reflect.ValueOf(elem).Elem().Set(reflect.ValueOf(ne.(types.NilableType).Type_()).Elem())
	}
	return err
}

func (c *Config) SearchRead(model string, args []interface{}, options interface{}, elem interface{}) error {
	ne := elem.(types.Type).NilableType_()
	err := c.DoRequest("search_read", model, args, options, ne)
	if err == nil {
		reflect.ValueOf(elem).Elem().Set(reflect.ValueOf(ne.(types.NilableType).Type_()).Elem())
	}
	return err
}

func (c *Config) SearchCount(model string, args []interface{}, elem interface{}) error {
	return c.DoRequest("search_count", model, args, nil, elem)
}

func (c *Config) DoRequest(method string, model string, args []interface{}, options interface{}, elem interface{}) error {
	client, err := c.NewClient()
	if err != nil {
		return err
	}
	return client.Call("execute_kw", []interface{}{c.Session.DbName, c.Session.UserID, c.Session.Password, model, method, args, options}, elem)
}

// Higher-level functions for data retrival
func (c *Config) getIdsByName(model string, name string) ([]int64, error) {
	var ids []int64
	err := c.Search(model, []interface{}{[]string{"name", "=", name}}, nil, &ids)
	return ids, err
}

func (c *Config) getByIds(model string, ids []int64, elem interface{}) error {
	err := c.Read(model, []interface{}{ids}, nil, elem)
	return err
}

func (c *Config) getByName(model string, name string, elem interface{}) error {
	err := c.SearchRead(model, []interface{}{[]interface{}{[]string{"name", "=", name}}}, nil, elem)
	return err
}

func (c *Config) getByField(model string, field string, value string, elem interface{}) error {
	err := c.SearchRead(model, []interface{}{[]interface{}{[]string{field, "=", value}}}, nil, elem)
	return err
}

func (c *Config) getAll(model string, elem interface{}) error {
	err := c.SearchRead(model, []interface{}{[]interface{}{}}, nil, elem)
	return err
}

// Higher-level functions for data manipulation
func (c *Config) create(model string, fields map[string]interface{}, relation *types.Relations) (int64, error) {
	var id int64
	if relation != nil {
		types.HandleRelations(&fields, relation)
	}
	err := c.Create(model, []interface{}{fields}, &id)
	return id, err
}

func (c *Config) update(model string, ids []int64, fields map[string]interface{}, relation *types.Relations) error {
	if relation != nil {
		types.HandleRelations(&fields, relation)
	}
	err := c.Update(model, []interface{}{ids, fields})
	return err
}

func (c *Config) delete(model string, ids []int64) error {
	return c.Delete(model, []interface{}{ids})
}
