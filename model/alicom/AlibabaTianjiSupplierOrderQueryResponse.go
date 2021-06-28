package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询供应商订单 APIResponse
alibaba.tianji.supplier.order.query

查询供应商订单
*/
type AlibabaTianjiSupplierOrderQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_tianji_supplier_order_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分销订单信息
    
    ModelList   []DistributionOrderInfo `json:"model_list,omitempty" xml:"