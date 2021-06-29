package tmallnr

// NrTimingFulfillDetailQueryRespDto 
type NrTimingFulfillDetailQueryRespDto struct {

    // 当前状态
    
    NrDeliveryBriefStatusDTO   *NrDeliveryBriefStatusDto `json:"nr_delivery_brief_status_d_t_o,omitempty" xml:"nr_delivery_brief_status_d_t_o,omitempty"`
    

    // 历史状态
    
    NrDeliveryBriefStatusDTOs   []NrDeliveryBriefStatusDto `json:"nr_delivery_brief_status_d_t_os,omitempty" xml:"nr_delivery_brief_status_d_t_os,omitempty"`
    

}
