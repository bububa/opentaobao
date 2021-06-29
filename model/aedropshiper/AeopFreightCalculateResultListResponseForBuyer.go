package aedropshiper

// AeopFreightCalculateResultListResponseForBuyer 
type AeopFreightCalculateResultListResponseForBuyer struct {

    // aeopFreightCalculateResultForBuyerDTOList
    
    AeopFreightCalculateResultForBuyerDTOList   []AeopFreightCalculateResultForBuyerDto `json:"aeop_freight_calculate_result_for_buyer_d_t_o_list,omitempty" xml:"aeop_freight_calculate_result_for_buyer_d_t_o_list,omitempty"`
    

    // errorDesc
    
    ErrorDesc   string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
