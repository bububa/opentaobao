package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询单据列表 APIResponse
taobao.wlb.wms.cainiao.bill.query

查询单据列表
*/
type TaobaoWlbWmsCainiaoBillQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbWmsCainiaoBillQueryResponse `json:"wlb_wms_cainiao_bill_query_response,omitempty"` 
    TaobaoWlbWmsCainiaoBillQueryResponse
}

/* model for simplify = false
type TaobaoWlbWmsCainiaoBillQueryResponse struct {

    // 总条数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 订单列表信息
    
    OrderInfoList  struct {
        CainiaoBillQueryOrderinfolist  []CainiaoBillQueryOrderinfolist `json:"cainiao_bill_query_orderinfolist,omitempty"`
    } `json:"order_info_list,omitempty"`
    

}
*/

type TaobaoWlbWmsCainiaoBillQueryResponse struct {

    // 总条数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 订单列表信息
    OrderInfoList   []CainiaoBillQueryOrderinfolist `json:"order_info_list,omitempty"`

}
