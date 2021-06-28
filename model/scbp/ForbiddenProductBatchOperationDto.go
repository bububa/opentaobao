package scbp

// ForbiddenProductBatchOperationDto 
/* model for simplify = false
type ForbiddenProductBatchOperationDto struct {

    // 查询条件
    
    ForbiddenProductOperationList  struct {
        ForbiddenProductOperationDto  []ForbiddenProductOperationDto `json:"forbidden_product_operation_dto,omitempty"`
    } `json:"forbidden_product_operation_list,omitempty"`
    

}
*/

// ForbiddenProductBatchOperationDto 
type ForbiddenProductBatchOperationDto struct {

    // 查询条件
    ForbiddenProductOperationList   []ForbiddenProductOperationDto `json:"forbidden_product_operation_list,omitempty"`

}
