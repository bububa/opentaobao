package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家取消获取的电子面单号v1.0 API返回值 
taobao.wlb.waybill.i.cancel

面单号有误需要取消的时候，调用该接口取消获取的电子面单。
*/
type TaobaoWlbWaybillICancelAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWaybillICancelResponse
}

// 商家取消获取的电子面单号v1.0 成功返回结果
type TaobaoWlbWaybillICancelResponse struct {
    XMLName xml.Name `xml:"wlb_waybill_i_cancel_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用取消是否成功
    CancelResult   bool `json:"cancel_result,omitempty" xml:"cancel_result,omitempty"`
}
