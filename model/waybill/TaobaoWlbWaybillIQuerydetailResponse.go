package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查面单号状态v1.0 APIResponse
taobao.wlb.waybill.i.querydetail

查看面单号的当前状态，如签收、发货、失效等。
*/
type TaobaoWlbWaybillIQuerydetailAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_waybill_i_querydetail_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 不存在的面单号
    
    InexistentWaybillCodes   []string `json:"inexistent_waybill_codes,omitempty" xml:"