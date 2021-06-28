package scbp

// AdGroupBatchOperationDto 
/* model for simplify = false
type AdGroupBatchOperationDto struct {

    // 入参
    
    AdGroupOperationList  struct {
        AdGroupOperationDto  []AdGroupOperationDto `json:"ad_group_operation_dto,omitempty"`
    } `json:"ad_group_operation_list,omitempty"`
    

}
*/

// AdGroupBatchOperationDto 
type AdGroupBatchOperationDto struct {

    // 入参
    AdGroupOperationList   []AdGroupOperationDto `json:"ad_group_operation_list,omitempty"`

}
