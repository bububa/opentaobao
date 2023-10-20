package youkuott

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttKittyCommonorderSyncAPIResponse 运营商一般订单同步 API返回值
// youku.ott.kitty.commonorder.sync
//
// 运营商一般订单同步
type YoukuOttKittyCommonorderSyncAPIResponse struct {
	model.CommonResponse
	YoukuOttKittyCommonorderSyncAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuOttKittyCommonorderSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuOttKittyCommonorderSyncAPIResponseModel).Reset()
}

// YoukuOttKittyCommonorderSyncAPIResponseModel is 运营商一般订单同步 成功返回结果
type YoukuOttKittyCommonorderSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_kitty_commonorder_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回码（一般情况请求方只需要关心success，除非特殊情况需要关心错误码）
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 是否成功 true:成功 false:失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *YoukuOttKittyCommonorderSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.RetCode = 0
	m.IsSuccess = false
}

var poolYoukuOttKittyCommonorderSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuOttKittyCommonorderSyncAPIResponse)
	},
}

// GetYoukuOttKittyCommonorderSyncAPIResponse 从 sync.Pool 获取 YoukuOttKittyCommonorderSyncAPIResponse
func GetYoukuOttKittyCommonorderSyncAPIResponse() *YoukuOttKittyCommonorderSyncAPIResponse {
	return poolYoukuOttKittyCommonorderSyncAPIResponse.Get().(*YoukuOttKittyCommonorderSyncAPIResponse)
}

// ReleaseYoukuOttKittyCommonorderSyncAPIResponse 将 YoukuOttKittyCommonorderSyncAPIResponse 保存到 sync.Pool
func ReleaseYoukuOttKittyCommonorderSyncAPIResponse(v *YoukuOttKittyCommonorderSyncAPIResponse) {
	v.Reset()
	poolYoukuOttKittyCommonorderSyncAPIResponse.Put(v)
}
