package pur

// MallReceivePrResponseData 
type MallReceivePrResponseData struct {

    // 采购商城申请单ID
    
    PurReqId   string `json:"pur_req_id,omitempty" xml:"pur_req_id,omitempty"`
    

    // 下单成功后跳转地址
    
    RedirectUrl   string `json:"redirect_url,omitempty" xml:"redirect_url,omitempty"`
    

}
