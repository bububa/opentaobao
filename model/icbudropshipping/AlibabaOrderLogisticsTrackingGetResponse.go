package icbudropshipping

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴订单物流轨迹查询 API返回值 
alibaba.order.logistics.tracking.get

阿里巴巴订单物流轨迹查询
*/
type AlibabaOrderLogisticsTrackingGetAPIResponse struct {
    model.CommonResponse
    AlibabaOrderLogisticsTrackingGetResponse
}

// 阿里巴巴订单物流轨迹查询 成功返回结果
type AlibabaOrderLogisticsTrackingGetResponse struct {
    XMLName xml.Name `xml:"alibaba_order_logistics_tracking_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // logistics  tracking List
    TrackingList   []LogisticsTracking `json:"tracking_list,omitempty" xml:"tracking_list>logistics_tracking,omitempty"`
}
