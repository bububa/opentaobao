package cmns

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosServiceCmnsCoaMessageCancelAPIResponse CMNS消息撤回 API返回值
// yunos.service.cmns.coa.message.cancel
//
// 此接口用户撤回之前已经发出去的消息，根据消息ID撤回，只能撤回此appKey创建的消息。
type YunosServiceCmnsCoaMessageCancelAPIResponse struct {
	model.CommonResponse
	YunosServiceCmnsCoaMessageCancelAPIResponseModel
}

// Reset 清空结构体
func (m *YunosServiceCmnsCoaMessageCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosServiceCmnsCoaMessageCancelAPIResponseModel).Reset()
}

// YunosServiceCmnsCoaMessageCancelAPIResponseModel is CMNS消息撤回 成功返回结果
type YunosServiceCmnsCoaMessageCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_service_cmns_coa_message_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用出错时返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回状态 200表示调用成功
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// true:撤回成功&lt;br/&gt;false:撤回失败
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *YunosServiceCmnsCoaMessageCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Status = 0
	m.Data = false
}

var poolYunosServiceCmnsCoaMessageCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosServiceCmnsCoaMessageCancelAPIResponse)
	},
}

// GetYunosServiceCmnsCoaMessageCancelAPIResponse 从 sync.Pool 获取 YunosServiceCmnsCoaMessageCancelAPIResponse
func GetYunosServiceCmnsCoaMessageCancelAPIResponse() *YunosServiceCmnsCoaMessageCancelAPIResponse {
	return poolYunosServiceCmnsCoaMessageCancelAPIResponse.Get().(*YunosServiceCmnsCoaMessageCancelAPIResponse)
}

// ReleaseYunosServiceCmnsCoaMessageCancelAPIResponse 将 YunosServiceCmnsCoaMessageCancelAPIResponse 保存到 sync.Pool
func ReleaseYunosServiceCmnsCoaMessageCancelAPIResponse(v *YunosServiceCmnsCoaMessageCancelAPIResponse) {
	v.Reset()
	poolYunosServiceCmnsCoaMessageCancelAPIResponse.Put(v)
}
