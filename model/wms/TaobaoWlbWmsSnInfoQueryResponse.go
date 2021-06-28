package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询单据序列号信息 APIResponse
taobao.wlb.wms.sn.info.query

查询仓库作业的各类单据记录的Sn信息
*/
type TaobaoWlbWmsSnInfoQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbWmsSnInfoQueryResponse `json:"wlb_wms_sn_info_query_response,omitempty"` 
    TaobaoWlbWmsSnInfoQueryResponse
}

/* model for simplify = false
type TaobaoWlbWmsSnInfoQueryResponse struct {

    // 接口返回
    
    Result  *struct {
        TaobaoWlbWmsSnInfoQueryResult  *TaobaoWlbWmsSnInfoQueryResult `json:"taobao_wlb_wms_sn_info_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWlbWmsSnInfoQueryResponse struct {

    // 接口返回
    Result   *TaobaoWlbWmsSnInfoQueryResult `json:"result,omitempty"`

}
