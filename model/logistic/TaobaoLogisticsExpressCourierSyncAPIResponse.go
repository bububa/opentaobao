package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressCourierSyncAPIResponse 快递公司同步小件员信息 API返回值
// taobao.logistics.express.courier.sync
//
// 快递公司同步小件员信息
type TaobaoLogisticsExpressCourierSyncAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressCourierSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressCourierSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressCourierSyncAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressCourierSyncAPIResponseModel is 快递公司同步小件员信息 成功返回结果
type TaobaoLogisticsExpressCourierSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_courier_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码描述
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 错误码标识
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 校验成功或者异常
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
	// 是否可重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressCourierSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizErrorMessage = ""
	m.BizErrorCode = ""
	m.Suc = false
	m.Retry = false
}

var poolTaobaoLogisticsExpressCourierSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressCourierSyncAPIResponse)
	},
}

// GetTaobaoLogisticsExpressCourierSyncAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressCourierSyncAPIResponse
func GetTaobaoLogisticsExpressCourierSyncAPIResponse() *TaobaoLogisticsExpressCourierSyncAPIResponse {
	return poolTaobaoLogisticsExpressCourierSyncAPIResponse.Get().(*TaobaoLogisticsExpressCourierSyncAPIResponse)
}

// ReleaseTaobaoLogisticsExpressCourierSyncAPIResponse 将 TaobaoLogisticsExpressCourierSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressCourierSyncAPIResponse(v *TaobaoLogisticsExpressCourierSyncAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressCourierSyncAPIResponse.Put(v)
}
