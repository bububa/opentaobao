package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品hscode信息审核状态查询接口 APIResponse
tmall.item.hscode.audit.results.query

通过此接口查询天猫跨境商品的hscode信息审核状态，卖家可以参考返回结果判断是否需要调整商品hscode相关信息。
*/
type TmallItemHscodeAuditResultsQueryAPIResponse struct {
    model.CommonResponse
    Response *TmallItemHscodeAuditResultsQueryResponse `json:"tmall_item_hscode_audit_results_query_response,omitempty"`
}

type TmallItemHscodeAuditResultsQueryResponse struct {

    // 商品或sku的hscode信息审核状态。
    ResultList   []HscodeAuditInfo `json:"result_list,omitempty"`

}
