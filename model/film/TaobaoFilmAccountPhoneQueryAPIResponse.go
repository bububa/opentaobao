package film

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmAccountPhoneQueryAPIResponse 根据手机查询匹配账号列表 API返回值
// taobao.film.account.phone.query
//
// 根据手机号查询匹配的账号列表
type TaobaoFilmAccountPhoneQueryAPIResponse struct {
	model.CommonResponse
	TaobaoFilmAccountPhoneQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFilmAccountPhoneQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFilmAccountPhoneQueryAPIResponseModel).Reset()
}

// TaobaoFilmAccountPhoneQueryAPIResponseModel is 根据手机查询匹配账号列表 成功返回结果
type TaobaoFilmAccountPhoneQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"film_account_phone_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultListModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFilmAccountPhoneQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFilmAccountPhoneQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFilmAccountPhoneQueryAPIResponse)
	},
}

// GetTaobaoFilmAccountPhoneQueryAPIResponse 从 sync.Pool 获取 TaobaoFilmAccountPhoneQueryAPIResponse
func GetTaobaoFilmAccountPhoneQueryAPIResponse() *TaobaoFilmAccountPhoneQueryAPIResponse {
	return poolTaobaoFilmAccountPhoneQueryAPIResponse.Get().(*TaobaoFilmAccountPhoneQueryAPIResponse)
}

// ReleaseTaobaoFilmAccountPhoneQueryAPIResponse 将 TaobaoFilmAccountPhoneQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFilmAccountPhoneQueryAPIResponse(v *TaobaoFilmAccountPhoneQueryAPIResponse) {
	v.Reset()
	poolTaobaoFilmAccountPhoneQueryAPIResponse.Put(v)
}
