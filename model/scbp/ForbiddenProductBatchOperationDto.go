package scbp

// ForbiddenProductBatchOperationDto 
type ForbiddenProductBatchOperationDto struct {
    // 查询条件
    ForbiddenProductOperationList   []ForbiddenProductOperationDto `json:"forbidden_product_operation_list,omitempty" xml:"forbidden_product_operation_list>forbidden_product_operation_dto,omitempty"`
}
