package einvoice

// InvoiceContactDTO 
type InvoiceContactDTO struct {
    // 联系人姓名
    ContactName   string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
    // 联系人电话
    ContactMobile   string `json:"contact_mobile,omitempty" xml:"contact_mobile,omitempty"`
    // 联系人地址
    ContactAddr   string `json:"contact_addr,omitempty" xml:"contact_addr,omitempty"`
    // 联系人邮件
    ContactMail   string `json:"contact_mail,omitempty" xml:"contact_mail,omitempty"`
}
