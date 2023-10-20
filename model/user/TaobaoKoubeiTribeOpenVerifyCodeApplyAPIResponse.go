package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse 口碑综合体手机号获取验证码 API返回值
// taobao.koubei.tribe.open.verify.code.apply
//
// 口碑综合体通过手机号获取验证码对外开放接口
type TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponseModel).Reset()
}

// TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponseModel is 口碑综合体手机号获取验证码 成功返回结果
type TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_tribe_open_verify_code_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoKoubeiTribeOpenVerifyCodeApplyResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse)
	},
}

// GetTaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse 从 sync.Pool 获取 TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse
func GetTaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse() *TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse {
	return poolTaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse.Get().(*TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse)
}

// ReleaseTaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse 将 TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse(v *TaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse) {
	v.Reset()
	poolTaobaoKoubeiTribeOpenVerifyCodeApplyAPIResponse.Put(v)
}
