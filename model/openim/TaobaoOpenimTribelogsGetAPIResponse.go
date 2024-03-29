package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribelogsGetAPIResponse openim 群聊天记录导出接口 API返回值
// taobao.openim.tribelogs.get
//
// 获取openim账号的群聊天记录
type TaobaoOpenimTribelogsGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribelogsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimTribelogsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimTribelogsGetAPIResponseModel).Reset()
}

// TaobaoOpenimTribelogsGetAPIResponseModel is openim 群聊天记录导出接口 成功返回结果
type TaobaoOpenimTribelogsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribelogs_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 错误码
	RetCode int64 `json:"retCode,omitempty" xml:"retCode,omitempty"`
	// 返回结构
	Data *TribeMessageResult `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimTribelogsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Reason = ""
	m.RetCode = 0
	m.Data = nil
	m.Succ = false
}

var poolTaobaoOpenimTribelogsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimTribelogsGetAPIResponse)
	},
}

// GetTaobaoOpenimTribelogsGetAPIResponse 从 sync.Pool 获取 TaobaoOpenimTribelogsGetAPIResponse
func GetTaobaoOpenimTribelogsGetAPIResponse() *TaobaoOpenimTribelogsGetAPIResponse {
	return poolTaobaoOpenimTribelogsGetAPIResponse.Get().(*TaobaoOpenimTribelogsGetAPIResponse)
}

// ReleaseTaobaoOpenimTribelogsGetAPIResponse 将 TaobaoOpenimTribelogsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimTribelogsGetAPIResponse(v *TaobaoOpenimTribelogsGetAPIResponse) {
	v.Reset()
	poolTaobaoOpenimTribelogsGetAPIResponse.Put(v)
}
