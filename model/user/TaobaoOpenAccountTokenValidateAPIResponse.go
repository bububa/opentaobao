package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountTokenValidateAPIResponse open account token验证 API返回值
// taobao.open.account.token.validate
//
// open account token验证
type TaobaoOpenAccountTokenValidateAPIResponse struct {
	model.CommonResponse
	TaobaoOpenAccountTokenValidateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenAccountTokenValidateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenAccountTokenValidateAPIResponseModel).Reset()
}

// TaobaoOpenAccountTokenValidateAPIResponseModel is open account token验证 成功返回结果
type TaobaoOpenAccountTokenValidateAPIResponseModel struct {
	XMLName xml.Name `xml:"open_account_token_validate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 验证成功返回token中的信息
	Data *OpenAccountTokenValidateResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenAccountTokenValidateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoOpenAccountTokenValidateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenAccountTokenValidateAPIResponse)
	},
}

// GetTaobaoOpenAccountTokenValidateAPIResponse 从 sync.Pool 获取 TaobaoOpenAccountTokenValidateAPIResponse
func GetTaobaoOpenAccountTokenValidateAPIResponse() *TaobaoOpenAccountTokenValidateAPIResponse {
	return poolTaobaoOpenAccountTokenValidateAPIResponse.Get().(*TaobaoOpenAccountTokenValidateAPIResponse)
}

// ReleaseTaobaoOpenAccountTokenValidateAPIResponse 将 TaobaoOpenAccountTokenValidateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenAccountTokenValidateAPIResponse(v *TaobaoOpenAccountTokenValidateAPIResponse) {
	v.Reset()
	poolTaobaoOpenAccountTokenValidateAPIResponse.Put(v)
}
