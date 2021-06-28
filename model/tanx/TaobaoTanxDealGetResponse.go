package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
对外部dsp提供交易id查询接口 APIResponse
taobao.tanx.deal.get

对外部dsp提供交易id查询接口
*/
type TaobaoTanxDealGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tanx_deal_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果代码
    
    Code   int64 `json:"code,omitempty" xml:"