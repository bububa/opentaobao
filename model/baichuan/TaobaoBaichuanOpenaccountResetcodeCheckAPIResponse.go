package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse
百川验证找回密码验证码 API返回值
taobao.baichuan.openaccount.resetcode.check

百川验证找回密码验证码 */
type TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountResetcodeCheckAPIResponseModel
}

// TaobaoBaichuanOpenaccountResetcodeCheckAPIResponseModel is 百川验证找回密码验证码 成功返回结果
type TaobaoBaichuanOpenaccountResetcodeCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_resetcode_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
