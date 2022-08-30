package tmallnr

// PageData 结构体
type PageData struct {
	// 电子凭证实例DTO
	DataList []NrtCertificateInstanceResponseDto `json:"data_list,omitempty" xml:"data_list>nrt_certificate_instance_response_dto,omitempty"`
	// 显示数量
	PageSize string `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页
	CurrentPage string `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总数
	TotalItem string `json:"total_item,omitempty" xml:"total_item,omitempty"`
}
