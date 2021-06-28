package mos

// CallDispatcherDto 
/* model for simplify = false
type CallDispatcherDto struct {

    // 包裹信息
    
    CodeInfoList  struct {
        CodeInfoDto  []CodeInfoDto `json:"code_info_dto,omitempty"`
    } `json:"code_info_list,omitempty"`
    

    // 主单号
    
    ParentNo   string `json:"parent_no,omitempty"`
    

    // 收货信息
    
    ReceiveInfo  *struct {
        DeliveryCustomDto  *DeliveryCustomDto `json:"delivery_custom_dto,omitempty"`
    } `json:"receive_info,omitempty"`
    

}
*/

// CallDispatcherDto 
type CallDispatcherDto struct {

    // 包裹信息
    CodeInfoList   []CodeInfoDto `json:"code_info_list,omitempty"`

    // 主单号
    ParentNo   string `json:"parent_no,omitempty"`

    // 收货信息
    ReceiveInfo   *DeliveryCustomDto `json:"receive_info,omitempty"`

}
