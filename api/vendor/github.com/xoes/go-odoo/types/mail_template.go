package types

import (
	"time"
)

type MailTemplate struct {
	LastUpdate          time.Time `xmlrpc:"__last_update"`
	AttachmentIds       []int64   `xmlrpc:"attachment_ids"`
	AutoDelete          bool      `xmlrpc:"auto_delete"`
	BodyHtml            string    `xmlrpc:"body_html"`
	Copyvalue           string    `xmlrpc:"copyvalue"`
	CreateDate          time.Time `xmlrpc:"create_date"`
	CreateUid           Many2One  `xmlrpc:"create_uid"`
	DisplayName         string    `xmlrpc:"display_name"`
	EmailCc             string    `xmlrpc:"email_cc"`
	EmailFrom           string    `xmlrpc:"email_from"`
	EmailTo             string    `xmlrpc:"email_to"`
	Id                  int64     `xmlrpc:"id"`
	Lang                string    `xmlrpc:"lang"`
	MailServerId        Many2One  `xmlrpc:"mail_server_id"`
	Model               string    `xmlrpc:"model"`
	ModelId             Many2One  `xmlrpc:"model_id"`
	ModelObjectField    Many2One  `xmlrpc:"model_object_field"`
	Name                string    `xmlrpc:"name"`
	NullValue           string    `xmlrpc:"null_value"`
	PartnerTo           string    `xmlrpc:"partner_to"`
	RefIrActWindow      Many2One  `xmlrpc:"ref_ir_act_window"`
	RefIrValue          Many2One  `xmlrpc:"ref_ir_value"`
	ReplyTo             string    `xmlrpc:"reply_to"`
	ReportName          string    `xmlrpc:"report_name"`
	ReportTemplate      Many2One  `xmlrpc:"report_template"`
	ScheduledDate       string    `xmlrpc:"scheduled_date"`
	SubModelObjectField Many2One  `xmlrpc:"sub_model_object_field"`
	SubObject           Many2One  `xmlrpc:"sub_object"`
	Subject             string    `xmlrpc:"subject"`
	UseDefaultTo        bool      `xmlrpc:"use_default_to"`
	UserSignature       bool      `xmlrpc:"user_signature"`
	WriteDate           time.Time `xmlrpc:"write_date"`
	WriteUid            Many2One  `xmlrpc:"write_uid"`
}

type MailTemplateNil struct {
	LastUpdate          interface{} `xmlrpc:"__last_update"`
	AttachmentIds       interface{} `xmlrpc:"attachment_ids"`
	AutoDelete          bool        `xmlrpc:"auto_delete"`
	BodyHtml            interface{} `xmlrpc:"body_html"`
	Copyvalue           interface{} `xmlrpc:"copyvalue"`
	CreateDate          interface{} `xmlrpc:"create_date"`
	CreateUid           interface{} `xmlrpc:"create_uid"`
	DisplayName         interface{} `xmlrpc:"display_name"`
	EmailCc             interface{} `xmlrpc:"email_cc"`
	EmailFrom           interface{} `xmlrpc:"email_from"`
	EmailTo             interface{} `xmlrpc:"email_to"`
	Id                  interface{} `xmlrpc:"id"`
	Lang                interface{} `xmlrpc:"lang"`
	MailServerId        interface{} `xmlrpc:"mail_server_id"`
	Model               interface{} `xmlrpc:"model"`
	ModelId             interface{} `xmlrpc:"model_id"`
	ModelObjectField    interface{} `xmlrpc:"model_object_field"`
	Name                interface{} `xmlrpc:"name"`
	NullValue           interface{} `xmlrpc:"null_value"`
	PartnerTo           interface{} `xmlrpc:"partner_to"`
	RefIrActWindow      interface{} `xmlrpc:"ref_ir_act_window"`
	RefIrValue          interface{} `xmlrpc:"ref_ir_value"`
	ReplyTo             interface{} `xmlrpc:"reply_to"`
	ReportName          interface{} `xmlrpc:"report_name"`
	ReportTemplate      interface{} `xmlrpc:"report_template"`
	ScheduledDate       interface{} `xmlrpc:"scheduled_date"`
	SubModelObjectField interface{} `xmlrpc:"sub_model_object_field"`
	SubObject           interface{} `xmlrpc:"sub_object"`
	Subject             interface{} `xmlrpc:"subject"`
	UseDefaultTo        bool        `xmlrpc:"use_default_to"`
	UserSignature       bool        `xmlrpc:"user_signature"`
	WriteDate           interface{} `xmlrpc:"write_date"`
	WriteUid            interface{} `xmlrpc:"write_uid"`
}

var MailTemplateModel string = "mail.template"

type MailTemplates []MailTemplate

type MailTemplatesNil []MailTemplateNil

func (s *MailTemplate) NilableType_() interface{} {
	return &MailTemplateNil{}
}

func (ns *MailTemplateNil) Type_() interface{} {
	s := &MailTemplate{}
	return load(ns, s)
}

func (s *MailTemplates) NilableType_() interface{} {
	return &MailTemplatesNil{}
}

func (ns *MailTemplatesNil) Type_() interface{} {
	s := &MailTemplates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailTemplate))
	}
	return s
}
