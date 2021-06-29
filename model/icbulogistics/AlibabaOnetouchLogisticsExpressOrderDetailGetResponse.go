package icbulogistics

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单详细信息(面单及仓库信息) APIResponse
alibaba.onetouch.logistics.express.order.detail.get

订单详细信息(面单及仓库信息)
*/
type AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaOnetouchLogisticsExpressOrderDetailGetResponse
}

type AlibabaOnetouchLogisticsExpressOrderDetailGetResponse struct {
    XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_order_detail_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaOnetouchLogisticsExpressOrderDetailGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
