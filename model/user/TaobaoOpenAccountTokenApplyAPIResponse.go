package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountTokenApplyAPIResponse 申请免登Open Account Token API返回值
// taobao.open.account.token.apply
//
// 申请免登Open Account Token
type TaobaoOpenAccountTokenApplyAPIResponse struct {
	model.CommonResponse
	TaobaoOpenAccountTokenApplyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenAccountTokenApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenAccountTokenApplyAPIResponseModel).Reset()
}

// TaobaoOpenAccountTokenApplyAPIResponseModel is 申请免登Open Account Token 成功返回结果
type TaobaoOpenAccountTokenApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"open_account_token_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的token结果
	Data *OpenAccountTokenApplyResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenAccountTokenApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoOpenAccountTokenApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenAccountTokenApplyAPIResponse)
	},
}

// GetTaobaoOpenAccountTokenApplyAPIResponse 从 sync.Pool 获取 TaobaoOpenAccountTokenApplyAPIResponse
func GetTaobaoOpenAccountTokenApplyAPIResponse() *TaobaoOpenAccountTokenApplyAPIResponse {
	return poolTaobaoOpenAccountTokenApplyAPIResponse.Get().(*TaobaoOpenAccountTokenApplyAPIResponse)
}

// ReleaseTaobaoOpenAccountTokenApplyAPIResponse 将 TaobaoOpenAccountTokenApplyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenAccountTokenApplyAPIResponse(v *TaobaoOpenAccountTokenApplyAPIResponse) {
	v.Reset()
	poolTaobaoOpenAccountTokenApplyAPIResponse.Put(v)
}
