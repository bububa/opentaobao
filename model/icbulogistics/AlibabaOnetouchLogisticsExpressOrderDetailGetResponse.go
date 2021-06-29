package icbulogistics

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单详细信息(面单及仓库信息) API返回值 
alibaba.onetouch.logistics.express.order.detail.get

订单详细信息(面单及仓库信息)
*/
type AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaOnetouchLogisticsExpressOrderDetailGetResponse
}

// 订单详细信息(面单及仓库信息) 成功返回结果
type AlibabaOnetouchLogisticsExpressOrderDetailGetResponse struct {
    XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_order_detail_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaOnetouchLogisticsExpressOrderDetailGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
