package drugtrace

// AlibabaAlihealthDrugKytQuerybatchprodResult 
type AlibabaAlihealthDrugKytQuerybatchprodResult struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 批次产品信息DTO
    Models   []BatchProductInfoDTO `json:"models,omitempty" xml:"models>batch_product_info_dto,omitempty"`
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // msgCode
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}
