package openim

import (
	"encoding/xml"

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

// TaobaoOpenimIoscertSandboxSetAPIResponseModel is 设置开发环境证书 成功返回结果
type TaobaoOpenimIoscertSandboxSetAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_ioscert_sandbox_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作成功
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}
