package aesolution

// ItemListResultDto 结构体
type ItemListResultDto struct {
	// error message
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// error code
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// total page
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// success or not
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// products total count
	ProductCount int64 `json:"product_count,omitempty" xml:"product_count,omitempty"`
	// error msg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// current page
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// product list
	AeopAEProductDisplayDTOList []ItemDisplayDto `json:"aeop_a_e_product_display_d_t_o_list,omitempty" xml:"aeop_a_e_product_display_d_t_o_list>item_display_dto,omitempty"`
}
