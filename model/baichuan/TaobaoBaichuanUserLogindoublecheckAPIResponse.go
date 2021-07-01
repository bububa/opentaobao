package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanUserLogindoublecheckAPIResponse
百川H5登录二次验证 API返回值
taobao.baichuan.user.logindoublecheck

百川H5登录二次验证 */
type TaobaoBaichuanUserLogindoublecheckAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanUserLogindoublecheckAPIResponseModel
}

// TaobaoBaichuanUserLogindoublecheckAPIResponseModel is 百川H5登录二次验证 成功返回结果
type TaobaoBaichuanUserLogindoublecheckAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_user_logindoublecheck_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
