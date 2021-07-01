package logistic

// WarehouseReverseUploadingDto 结构体
type WarehouseReverseUploadingDto struct {
	// 销退单列表
	UploadingReverseDTOList []UploadingReverseDto `json:"uploading_reverse_d_t_o_list,omitempty" xml:"uploading_reverse_d_t_o_list>uploading_reverse_dto,omitempty"`
}
