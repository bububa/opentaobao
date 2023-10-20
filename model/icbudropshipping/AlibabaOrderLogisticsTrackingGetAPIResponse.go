package icbudropshipping

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOrderLogisticsTrackingGetAPIResponse 阿里巴巴订单物流轨迹查询 API返回值
// alibaba.order.logistics.tracking.get
//
// 阿里巴巴订单物流轨迹查询
type AlibabaOrderLogisticsTrackingGetAPIResponse struct {
	model.CommonResponse
	AlibabaOrderLogisticsTrackingGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOrderLogisticsTrackingGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOrderLogisticsTrackingGetAPIResponseModel).Reset()
}

// AlibabaOrderLogisticsTrackingGetAPIResponseModel is 阿里巴巴订单物流轨迹查询 成功返回结果
type AlibabaOrderLogisticsTrackingGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_order_logistics_tracking_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// logistics  tracking List
	TrackingList []LogisticsTracking `json:"tracking_list,omitempty" xml:"tracking_list>logistics_tracking,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOrderLogisticsTrackingGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TrackingList = m.TrackingList[:0]
}

var poolAlibabaOrderLogisticsTrackingGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOrderLogisticsTrackingGetAPIResponse)
	},
}

// GetAlibabaOrderLogisticsTrackingGetAPIResponse 从 sync.Pool 获取 AlibabaOrderLogisticsTrackingGetAPIResponse
func GetAlibabaOrderLogisticsTrackingGetAPIResponse() *AlibabaOrderLogisticsTrackingGetAPIResponse {
	return poolAlibabaOrderLogisticsTrackingGetAPIResponse.Get().(*AlibabaOrderLogisticsTrackingGetAPIResponse)
}

// ReleaseAlibabaOrderLogisticsTrackingGetAPIResponse 将 AlibabaOrderLogisticsTrackingGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOrderLogisticsTrackingGetAPIResponse(v *AlibabaOrderLogisticsTrackingGetAPIResponse) {
	v.Reset()
	poolAlibabaOrderLogisticsTrackingGetAPIResponse.Put(v)
}
