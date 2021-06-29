package larkiot

// BizSingleResult 
type BizSingleResult struct {
    // 业务订结果
    BizMsg   string `json:"biz_msg,omitempty" xml:"biz_msg,omitempty"`
    // 业务码
    BizCode   string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
    // 接口是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 业务返回
    Data   *IotGoodsOrderRsp `json:"data,omitempty" xml:"data,omitempty"`
}
