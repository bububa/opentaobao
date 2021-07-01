package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse
百川检查注册验证码 API返回值
taobao.baichuan.openaccount.registercode.check

百川检查注册验证码 */
type TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponseModel
}

// TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponseModel is 百川检查注册验证码 成功返回结果
type TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_registercode_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
