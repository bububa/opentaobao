package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse 上门取退运费支付状态查询接口 API返回值
// taobao.logistics.express.order.pay.tms.query
//
// 上门取退运费支付状态查询接口
type TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressOrderPayTmsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressOrderPayTmsQueryAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressOrderPayTmsQueryAPIResponseModel is 上门取退运费支付状态查询接口 成功返回结果
type TaobaoLogisticsExpressOrderPayTmsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_order_pay_tms_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 错误码
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 业务处理结果
	Data *Tms2MscPayQueryResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 是否可重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
	// 是	系统成功失败 true|false
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressOrderPayTmsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizErrorMessage = ""
	m.BizErrorCode = ""
	m.Data = nil
	m.Retry = false
	m.Suc = false
}

var poolTaobaoLogisticsExpressOrderPayTmsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse)
	},
}

// GetTaobaoLogisticsExpressOrderPayTmsQueryAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse
func GetTaobaoLogisticsExpressOrderPayTmsQueryAPIResponse() *TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse {
	return poolTaobaoLogisticsExpressOrderPayTmsQueryAPIResponse.Get().(*TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse)
}

// ReleaseTaobaoLogisticsExpressOrderPayTmsQueryAPIResponse 将 TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressOrderPayTmsQueryAPIResponse(v *TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressOrderPayTmsQueryAPIResponse.Put(v)
}
