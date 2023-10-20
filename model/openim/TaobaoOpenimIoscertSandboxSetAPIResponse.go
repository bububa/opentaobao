package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimIoscertSandboxSetAPIResponse 设置开发环境证书 API返回值
// taobao.openim.ioscert.sandbox.set
//
// 设置开发环境证书
type TaobaoOpenimIoscertSandboxSetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimIoscertSandboxSetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimIoscertSandboxSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimIoscertSandboxSetAPIResponseModel).Reset()
}

// TaobaoOpenimIoscertSandboxSetAPIResponseModel is 设置开发环境证书 成功返回结果
type TaobaoOpenimIoscertSandboxSetAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_ioscert_sandbox_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作成功
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimIoscertSandboxSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Code = ""
}

var poolTaobaoOpenimIoscertSandboxSetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimIoscertSandboxSetAPIResponse)
	},
}

// GetTaobaoOpenimIoscertSandboxSetAPIResponse 从 sync.Pool 获取 TaobaoOpenimIoscertSandboxSetAPIResponse
func GetTaobaoOpenimIoscertSandboxSetAPIResponse() *TaobaoOpenimIoscertSandboxSetAPIResponse {
	return poolTaobaoOpenimIoscertSandboxSetAPIResponse.Get().(*TaobaoOpenimIoscertSandboxSetAPIResponse)
}

// ReleaseTaobaoOpenimIoscertSandboxSetAPIResponse 将 TaobaoOpenimIoscertSandboxSetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimIoscertSandboxSetAPIResponse(v *TaobaoOpenimIoscertSandboxSetAPIResponse) {
	v.Reset()
	poolTaobaoOpenimIoscertSandboxSetAPIResponse.Put(v)
}
