package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取物流服务商电子面单号v1.0 APIResponse
taobao.wlb.waybill.i.get

商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。
*/
type TaobaoWlbWaybillIGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_waybill_i_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 面单申请接口返回信息
    
    WaybillApplyNewCols   []WaybillApplyNewInfo `json:"waybill_apply_new_cols,omitempty" xml:"