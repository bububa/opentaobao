package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家取消获取的电子面单号 APIResponse
cainiao.waybill.ii.cancel

面单号有误需要取消的时候，调用该接口取消获取的电子面单。
*/
type CainiaoWaybillIiCancelAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_waybill_ii_cancel_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用取消是否成功
    
    CancelResult   bool `json:"cancel_result,omitempty" xml:"