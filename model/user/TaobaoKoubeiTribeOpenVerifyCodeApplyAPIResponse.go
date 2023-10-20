package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaokoubeitribeopenverifycodeapplyAPIResponse 口碑综合体手机号获取验证码 API返回值
// taobao.koubei.tribe.open.verify.code.apply
//
// 口碑综合体通过手机号获取验证码对外开放接口
type TaobaokoubeitribeopenverifycodeapplyAPIResponse struct {
	model.CommonResponse
	TaobaokoubeitribeopenverifycodeapplyAPIResponseModel
}

// TaobaokoubeitribeopenverifycodeapplyAPIResponseModel is 口碑综合体手机号获取验证码 成功返回结果
type TaobaokoubeitribeopenverifycodeapplyAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_tribe_open_verify_code_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaokoubeitribeopenverifycodeapplyResult `json:"result,omitempty" xml:"result,omitempty"`
}
