package omniorder

// GetStoreConsignCodeResponse 
type GetStoreConsignCodeResponse struct {
    // 面单号
    MailNo   string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
    // 包裹Id, 后续发货需要使用
    PackageId   string `json:"package_id,omitempty" xml:"package_id,omitempty"`
    // 打印机内容
    PrintData   string `json:"print_data,omitempty" xml:"print_data,omitempty"`
    // 菜鸟生成的标签号
    TagCode   string `json:"tag_code,omitempty" xml:"tag_code,omitempty"`
}
