package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappEleuserPhoneGetAPIResponse
获取饿了么用户信息 API返回值
taobao.miniapp.eleuser.phone.get

获取饿了么用户信息 */
type TaobaoMiniappEleuserPhoneGetAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappEleuserPhoneGetAPIResponseModel
}

// TaobaoMiniappEleuserPhoneGetAPIResponseModel is 获取饿了么用户信息 成功返回结果
type TaobaoMiniappEleuserPhoneGetAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_eleuser_phone_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *EleUicInfo `json:"result,omitempty" xml:"result,omitempty"`
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
}
