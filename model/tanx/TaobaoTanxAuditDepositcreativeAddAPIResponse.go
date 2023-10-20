package tanx

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotanxauditdepositcreativeaddAPIResponse dsp托管创意新增接口 API返回值
// taobao.tanx.audit.depositcreative.add
//
// dsp托管创意新增接口
type TaobaotanxauditdepositcreativeaddAPIResponse struct {
	model.CommonResponse
	TaobaotanxauditdepositcreativeaddAPIResponseModel
}

// TaobaotanxauditdepositcreativeaddAPIResponseModel is dsp托管创意新增接口 成功返回结果
type TaobaotanxauditdepositcreativeaddAPIResponseModel struct {
	XMLName xml.Name `xml:"tanx_audit_depositcreative_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用的成功信息或失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用返回码
	TanxErrorCode int64 `json:"tanx_error_code,omitempty" xml:"tanx_error_code,omitempty"`
	// 是否成功
	IsOk bool `json:"is_ok,omitempty" xml:"is_ok,omitempty"`
}
