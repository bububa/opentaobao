package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressServiceSyncAPIResponse 服务信息回告接口 API返回值
// taobao.logistics.express.service.sync
//
// 服务信息回告接口
type TaobaoLogisticsExpressServiceSyncAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressServiceSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressServiceSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressServiceSyncAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressServiceSyncAPIResponseModel is 服务信息回告接口 成功返回结果
type TaobaoLogisticsExpressServiceSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_service_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码描述
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 校验成功或者异常
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
	// 错误码标识
	BizErrorCode bool `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 是否可重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressServiceSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizErrorMessage = ""
	m.Suc = false
	m.BizErrorCode = false
	m.Retry = false
}

var poolTaobaoLogisticsExpressServiceSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressServiceSyncAPIResponse)
	},
}

// GetTaobaoLogisticsExpressServiceSyncAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressServiceSyncAPIResponse
func GetTaobaoLogisticsExpressServiceSyncAPIResponse() *TaobaoLogisticsExpressServiceSyncAPIResponse {
	return poolTaobaoLogisticsExpressServiceSyncAPIResponse.Get().(*TaobaoLogisticsExpressServiceSyncAPIResponse)
}

// ReleaseTaobaoLogisticsExpressServiceSyncAPIResponse 将 TaobaoLogisticsExpressServiceSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressServiceSyncAPIResponse(v *TaobaoLogisticsExpressServiceSyncAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressServiceSyncAPIResponse.Put(v)
}
