package logistic

// WarehouseReverseUploadingDto 
/* model for simplify = false
type WarehouseReverseUploadingDto struct {

    // 销退单列表
    
    UploadingReverseDTOList  struct {
        UploadingReverseDto  []UploadingReverseDto `json:"uploading_reverse_dto,omitempty"`
    } `json:"uploading_reverse_d_t_o_list,omitempty"`
    

}
*/

// WarehouseReverseUploadingDto 
type WarehouseReverseUploadingDto struct {

    // 销退单列表
    UploadingReverseDTOList   []UploadingReverseDto `json:"uploading_reverse_d_t_o_list,omitempty"`

}
