package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressCollectSyncAPIResponse 服饰逆向揽收信息同步 API返回值
// taobao.logistics.express.collect.sync
//
// 服饰逆向揽收信息同步
type TaobaoLogisticsExpressCollectSyncAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressCollectSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressCollectSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressCollectSyncAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressCollectSyncAPIResponseModel is 服饰逆向揽收信息同步 成功返回结果
type TaobaoLogisticsExpressCollectSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_collect_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码描述
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 错误码标识
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 返回值
	Data *TmsCollectResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 校验成功或者异常
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
	// 是否可重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressCollectSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizErrorMessage = ""
	m.BizErrorCode = ""
	m.Data = nil
	m.Suc = false
	m.Retry = false
}

var poolTaobaoLogisticsExpressCollectSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressCollectSyncAPIResponse)
	},
}

// GetTaobaoLogisticsExpressCollectSyncAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressCollectSyncAPIResponse
func GetTaobaoLogisticsExpressCollectSyncAPIResponse() *TaobaoLogisticsExpressCollectSyncAPIResponse {
	return poolTaobaoLogisticsExpressCollectSyncAPIResponse.Get().(*TaobaoLogisticsExpressCollectSyncAPIResponse)
}

// ReleaseTaobaoLogisticsExpressCollectSyncAPIResponse 将 TaobaoLogisticsExpressCollectSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressCollectSyncAPIResponse(v *TaobaoLogisticsExpressCollectSyncAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressCollectSyncAPIResponse.Put(v)
}
