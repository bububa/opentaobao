package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
负卖库存失效接口 APIResponse
alibaba.ascp.aic.supplier.aicinventory.negative.sale.invalidate

失效负卖库存数据
*/
type AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIResponse struct {
    model.CommonResponse
    AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateResponse
}

type AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_aic_supplier_aicinventory_negative_sale_invalidate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值包装,result为返回具体消息内容
    
    FutureInvItemResponse   *ResultWrapper `json:"future_inv_item_response,omitempty" xml:"future_inv_item_response,omitempty"`

    
}
