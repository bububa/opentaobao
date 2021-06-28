package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营产品投放 APIResponse
taobao.alitrip.totoro.auxproduct.push

廉航辅营产品投放接口
*/
type TaobaoAlitripTotoroAuxproductPushAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_totoro_auxproduct_push_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作日志id，商家可通过该id在后台查看本次操作的具体结果
    
    TracerId   string `json:"tracer_id,omitempty" xml:"