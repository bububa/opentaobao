package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressDeliveryCutNotifyAPIResponse TMS配拦截结果回告 API返回值
// taobao.logistics.express.delivery.cut.notify
//
// TMS配拦截结果回告
type TaobaoLogisticsExpressDeliveryCutNotifyAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressDeliveryCutNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressDeliveryCutNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressDeliveryCutNotifyAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressDeliveryCutNotifyAPIResponseModel is TMS配拦截结果回告 成功返回结果
type TaobaoLogisticsExpressDeliveryCutNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_delivery_cut_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 是否成功
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
	// 错误编码
	BizErrorCode bool `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 是否支持重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressDeliveryCutNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizErrorMessage = ""
	m.Suc = false
	m.BizErrorCode = false
	m.Retry = false
}

var poolTaobaoLogisticsExpressDeliveryCutNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressDeliveryCutNotifyAPIResponse)
	},
}

// GetTaobaoLogisticsExpressDeliveryCutNotifyAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressDeliveryCutNotifyAPIResponse
func GetTaobaoLogisticsExpressDeliveryCutNotifyAPIResponse() *TaobaoLogisticsExpressDeliveryCutNotifyAPIResponse {
	return poolTaobaoLogisticsExpressDeliveryCutNotifyAPIResponse.Get().(*TaobaoLogisticsExpressDeliveryCutNotifyAPIResponse)
}

// ReleaseTaobaoLogisticsExpressDeliveryCutNotifyAPIResponse 将 TaobaoLogisticsExpressDeliveryCutNotifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressDeliveryCutNotifyAPIResponse(v *TaobaoLogisticsExpressDeliveryCutNotifyAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressDeliveryCutNotifyAPIResponse.Put(v)
}
