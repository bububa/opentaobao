package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AIC负卖库存新增和修改接口 APIResponse
alibaba.ascp.aic.supplier.aicinventory.negative.sale.publish

新增负卖库存记录和变更负卖库存记录
*/
type AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIResponse struct {
    model.CommonResponse
    AlibabaAscpAicSupplierAicinventoryNegativeSalePublishResponse
}

type AlibabaAscpAicSupplierAicinventoryNegativeSalePublishResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_aic_supplier_aicinventory_negative_sale_publish_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值包装,result为返回具体消息内容
    
    FutureInvItemResponse   *ResultWrapper `json:"future_inv_item_response,omitempty" xml:"future_inv_item_response,omitempty"`

    
}
