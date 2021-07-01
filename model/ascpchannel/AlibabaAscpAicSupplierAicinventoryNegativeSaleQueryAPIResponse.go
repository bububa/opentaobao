package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家查询负卖库存 API返回值 
alibaba.ascp.aic.supplier.aicinventory.negative.sale.query

商家根据当前接口查询负卖货品的库存
*/
type AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponseModel
}

// 商家查询负卖库存 成功返回结果
type AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ascp_aic_supplier_aicinventory_negative_sale_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值包装,result为返回具体消息内容
    AicinventoryQueryResponse   *ResultWrapper `json:"aicinventory_query_response,omitempty" xml:"aicinventory_query_response,omitempty"`
}
