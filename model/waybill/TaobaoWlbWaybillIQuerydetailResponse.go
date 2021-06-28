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
    // Response *TaobaoWlbWaybillIQuerydetailResponse `json:"wlb_waybill_i_querydetail_response,omitempty"` 
    TaobaoWlbWaybillIQuerydetailResponse
}

/* model for simplify = false
type TaobaoWlbWaybillIQuerydetailResponse struct {

    // 不存在的面单号
    
    InexistentWaybillCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"inexistent_waybill_codes,omitempty"`
    

    // 查询是否成功
    
    QuerySuccess   bool `json:"query_success,omitempty"`
    

    // 面单详情
    
    WaybillDetails  struct {
        WaybillDetailQueryInfo  []WaybillDetailQueryInfo `json:"waybill_detail_query_info,omitempty"`
    } `json:"waybill_details,omitempty"`
    

    // 面单查询错误编码
    
    ErrorCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"error_codes,omitempty"`
    

}
*/

type TaobaoWlbWaybillIQuerydetailResponse struct {

    // 不存在的面单号
    InexistentWaybillCodes   []string `json:"inexistent_waybill_codes,omitempty"`

    // 查询是否成功
    QuerySuccess   bool `json:"query_success,omitempty"`

    // 面单详情
    WaybillDetails   []WaybillDetailQueryInfo `json:"waybill_details,omitempty"`

    // 面单查询错误编码
    ErrorCodes   []string `json:"error_codes,omitempty"`

}
