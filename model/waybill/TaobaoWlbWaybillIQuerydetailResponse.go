package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查面单号状态v1.0 APIResponse
taobao.wlb.waybill.i.querydetail

查看面单号的当前状态，如签收、发货、失效等。
*/
type TaobaoWlbWaybillIQuerydetailAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbWaybillIQuerydetailResponse `json:"taobao_wlb_waybill_i_querydetail_response,omitempty"`
}

type TaobaoWlbWaybillIQuerydetailResponse struct {

    // 不存在的面单号
    InexistentWaybillCodes   []String `json:"inexistent_waybill_codes,omitempty"`

    // 查询是否成功
    QuerySuccess   bool `json:"query_success,omitempty"`

    // 面单详情
    WaybillDetails   []WaybillDetailQueryInfo `json:"waybill_details,omitempty"`

    // 面单查询错误编码
    ErrorCodes   []String `json:"error_codes,omitempty"`

}
