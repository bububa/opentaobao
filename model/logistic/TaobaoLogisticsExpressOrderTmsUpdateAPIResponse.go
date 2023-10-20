package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressOrderTmsUpdateAPIResponse 服务商修改上门取退时间接口 API返回值
// taobao.logistics.express.order.tms.update
//
// 服务商修改上门取退时间接口
type TaobaoLogisticsExpressOrderTmsUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressOrderTmsUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressOrderTmsUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressOrderTmsUpdateAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressOrderTmsUpdateAPIResponseModel is 服务商修改上门取退时间接口 成功返回结果
type TaobaoLogisticsExpressOrderTmsUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_order_tms_update_response"`
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
func (m *TaobaoLogisticsExpressOrderTmsUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizErrorMessage = ""
	m.BizErrorCode = ""
	m.Suc = false
	m.Retry = false
}

var poolTaobaoLogisticsExpressOrderTmsUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressOrderTmsUpdateAPIResponse)
	},
}

// GetTaobaoLogisticsExpressOrderTmsUpdateAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressOrderTmsUpdateAPIResponse
func GetTaobaoLogisticsExpressOrderTmsUpdateAPIResponse() *TaobaoLogisticsExpressOrderTmsUpdateAPIResponse {
	return poolTaobaoLogisticsExpressOrderTmsUpdateAPIResponse.Get().(*TaobaoLogisticsExpressOrderTmsUpdateAPIResponse)
}

// ReleaseTaobaoLogisticsExpressOrderTmsUpdateAPIResponse 将 TaobaoLogisticsExpressOrderTmsUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressOrderTmsUpdateAPIResponse(v *TaobaoLogisticsExpressOrderTmsUpdateAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressOrderTmsUpdateAPIResponse.Put(v)
}
