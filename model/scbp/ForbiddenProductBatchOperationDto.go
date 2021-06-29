package scbp

// ForbiddenProductBatchOperationDTO 
type ForbiddenProductBatchOperationDTO struct {
    // 查询条件
    ForbiddenProductOperationList   []ForbiddenProductOperationDTO `json:"forbidden_product_operation_list,omitempty" xml:"forbidden_product_operation_list>forbidden_product_operation_dto,omitempty"`
}
