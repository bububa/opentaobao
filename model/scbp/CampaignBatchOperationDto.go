package scbp

// CampaignBatchOperationDto 
/* model for simplify = false
type CampaignBatchOperationDto struct {

    // 具体操作实例
    
    CampaignOperationList  struct {
        Campaignoperationlist  []Campaignoperationlist `json:"campaignoperationlist,omitempty"`
    } `json:"campaign_operation_list,omitempty"`
    

}
*/

// CampaignBatchOperationDto 
type CampaignBatchOperationDto struct {

    // 具体操作实例
    CampaignOperationList   []Campaignoperationlist `json:"campaign_operation_list,omitempty"`

}
