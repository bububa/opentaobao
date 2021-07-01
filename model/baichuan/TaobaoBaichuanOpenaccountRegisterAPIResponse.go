package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountRegisterAPIResponse
百川账号注册 API返回值
taobao.baichuan.openaccount.register

百川账号注册 */
type TaobaoBaichuanOpenaccountRegisterAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountRegisterAPIResponseModel
}

// TaobaoBaichuanOpenaccountRegisterAPIResponseModel is 百川账号注册 成功返回结果
type TaobaoBaichuanOpenaccountRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
