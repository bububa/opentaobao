package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressOrderTmsCancelAPIResponse 服务商上门取退时间取消接口 API返回值
// taobao.logistics.express.order.tms.cancel
//
// 服务商上门取退时间取消接口
type TaobaoLogisticsExpressOrderTmsCancelAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressOrderTmsCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressOrderTmsCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressOrderTmsCancelAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressOrderTmsCancelAPIResponseModel is 服务商上门取退时间取消接口 成功返回结果
type TaobaoLogisticsExpressOrderTmsCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_order_tms_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 响应码
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 系统成功失败
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
	// 是否可重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressOrderTmsCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizErrorMessage = ""
	m.BizErrorCode = ""
	m.Suc = false
	m.Retry = false
}

var poolTaobaoLogisticsExpressOrderTmsCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressOrderTmsCancelAPIResponse)
	},
}

// GetTaobaoLogisticsExpressOrderTmsCancelAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressOrderTmsCancelAPIResponse
func GetTaobaoLogisticsExpressOrderTmsCancelAPIResponse() *TaobaoLogisticsExpressOrderTmsCancelAPIResponse {
	return poolTaobaoLogisticsExpressOrderTmsCancelAPIResponse.Get().(*TaobaoLogisticsExpressOrderTmsCancelAPIResponse)
}

// ReleaseTaobaoLogisticsExpressOrderTmsCancelAPIResponse 将 TaobaoLogisticsExpressOrderTmsCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressOrderTmsCancelAPIResponse(v *TaobaoLogisticsExpressOrderTmsCancelAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressOrderTmsCancelAPIResponse.Put(v)
}
