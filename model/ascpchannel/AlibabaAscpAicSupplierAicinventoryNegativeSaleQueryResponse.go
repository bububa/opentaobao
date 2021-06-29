package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家查询负卖库存 APIResponse
alibaba.ascp.aic.supplier.aicinventory.negative.sale.query

商家根据当前接口查询负卖货品的库存
*/
type AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryResponse
}

type AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_aic_supplier_aicinventory_negative_sale_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值包装,result为返回具体消息内容
    
    AicinventoryQueryResponse   *ResultWrapper `json:"aicinventory_query_response,omitempty" xml:"aicinventory_query_response,omitempty"`

    
}
