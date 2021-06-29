package icbudropshipping

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴订单物流轨迹查询 APIResponse
alibaba.order.logistics.tracking.get

阿里巴巴订单物流轨迹查询
*/
type AlibabaOrderLogisticsTrackingGetAPIResponse struct {
    model.CommonResponse
    AlibabaOrderLogisticsTrackingGetResponse
}

type AlibabaOrderLogisticsTrackingGetResponse struct {
    XMLName xml.Name `xml:"alibaba_order_logistics_tracking_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // logistics  tracking List
    
    TrackingList   []LogisticsTracking `json:"tracking_list,omitempty" xml:"tracking_list>logistics_tracking,omitempty"`
    
    
}
