package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询供应商订单 API返回值 
alibaba.tianji.supplier.order.query

查询供应商订单
*/
type AlibabaTianjiSupplierOrderQueryAPIResponse struct {
    model.CommonResponse
    AlibabaTianjiSupplierOrderQueryAPIResponseModel
}

// 查询供应商订单 成功返回结果
type AlibabaTianjiSupplierOrderQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_tianji_supplier_order_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 分销订单信息
    ModelList   []DistributionOrderInfo `json:"model_list,omitempty" xml:"model_list>distribution_order_info,omitempty"`
    // 查询总数
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
