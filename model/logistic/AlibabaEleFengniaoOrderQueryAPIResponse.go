package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoOrderQueryAPIResponse 查询订单基本信息 API返回值
// alibaba.ele.fengniao.order.query
//
// 查询订单基本信息
type AlibabaEleFengniaoOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaEleFengniaoOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleFengniaoOrderQueryAPIResponseModel).Reset()
}

// AlibabaEleFengniaoOrderQueryAPIResponseModel is 查询订单基本信息 成功返回结果
type AlibabaEleFengniaoOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_fengniao_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// kvs
	Kvs []Kvs `json:"kvs,omitempty" xml:"kvs>kvs,omitempty"`
	// 收件人经度
	ReceiverLongitude string `json:"receiver_longitude,omitempty" xml:"receiver_longitude,omitempty"`
	// 寄件人纬度
	TransportLatitude string `json:"transport_latitude,omitempty" xml:"transport_latitude,omitempty"`
	// 寄件人经度
	TransportLongitude string `json:"transport_longitude,omitempty" xml:"transport_longitude,omitempty"`
	// 收件人纬度
	ReceiverLatitude string `json:"receiver_latitude,omitempty" xml:"receiver_latitude,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Kvs = m.Kvs[:0]
	m.ReceiverLongitude = ""
	m.TransportLatitude = ""
	m.TransportLongitude = ""
	m.ReceiverLatitude = ""
	m.OrderId = 0
}

var poolAlibabaEleFengniaoOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleFengniaoOrderQueryAPIResponse)
	},
}

// GetAlibabaEleFengniaoOrderQueryAPIResponse 从 sync.Pool 获取 AlibabaEleFengniaoOrderQueryAPIResponse
func GetAlibabaEleFengniaoOrderQueryAPIResponse() *AlibabaEleFengniaoOrderQueryAPIResponse {
	return poolAlibabaEleFengniaoOrderQueryAPIResponse.Get().(*AlibabaEleFengniaoOrderQueryAPIResponse)
}

// ReleaseAlibabaEleFengniaoOrderQueryAPIResponse 将 AlibabaEleFengniaoOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleFengniaoOrderQueryAPIResponse(v *AlibabaEleFengniaoOrderQueryAPIResponse) {
	v.Reset()
	poolAlibabaEleFengniaoOrderQueryAPIResponse.Put(v)
}
