package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询供应商订单 APIResponse
alibaba.tianji.supplier.order.query

查询供应商订单
*/
type AlibabaTianjiSupplierOrderQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaTianjiSupplierOrderQueryResponse `json:"alibaba_tianji_supplier_order_query_response,omitempty"`
}

type AlibabaTianjiSupplierOrderQueryResponse struct {

    // 分销订单信息
    ModelList   []DistributionOrderInfo `json:"model_list,omitempty"`

    // 查询总数
    TotalCount   int64 `json:"total_count,omitempty"`

}
