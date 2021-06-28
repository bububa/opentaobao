package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营产品投放 APIResponse
taobao.alitrip.totoro.auxproduct.push

廉航辅营产品投放接口
*/
type TaobaoAlitripTotoroAuxproductPushAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripTotoroAuxproductPushResponse `json:"alitrip_totoro_auxproduct_push_response,omitempty"` 
    TaobaoAlitripTotoroAuxproductPushResponse
}

/* model for simplify = false
type TaobaoAlitripTotoroAuxproductPushResponse struct {

    // 操作日志id，商家可通过该id在后台查看本次操作的具体结果
    
    TracerId   string `json:"tracer_id,omitempty"`
    

    // 备注
    
    Message   string `json:"message,omitempty"`
    

    // 是否操作成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoAlitripTotoroAuxproductPushResponse struct {

    // 操作日志id，商家可通过该id在后台查看本次操作的具体结果
    TracerId   string `json:"tracer_id,omitempty"`

    // 备注
    Message   string `json:"message,omitempty"`

    // 是否操作成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
