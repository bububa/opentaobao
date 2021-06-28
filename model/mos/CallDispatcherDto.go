package mos

// CallDispatcherDto 
type CallDispatcherDto struct {

    // 包裹信息
    
    CodeInfoList   []CodeInfoDto `json:"code_info_list,omitempty" xml:"code_info_list,omitempty"`
    

    // 主单号
    
    ParentNo   string `json:"parent_no,omitempty" xml:"parent_no,omitempty"`
    

    // 收货信息
    
    ReceiveInfo   *DeliveryCustomDto `json:"receive_info,omitempty" xml:"receive_info,omitempty"`
    

}
