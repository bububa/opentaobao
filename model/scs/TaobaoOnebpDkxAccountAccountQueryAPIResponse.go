package scs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxAccountAccountQueryAPIResponse 万相台账号余额查询 API返回值
// taobao.onebp.dkx.account.account.query
//
// 万相台账号余额查询
type TaobaoOnebpDkxAccountAccountQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxAccountAccountQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxAccountAccountQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxAccountAccountQueryAPIResponseModel).Reset()
}

// TaobaoOnebpDkxAccountAccountQueryAPIResponseModel is 万相台账号余额查询 成功返回结果
type TaobaoOnebpDkxAccountAccountQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_account_account_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxAccountAccountQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxAccountAccountQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxAccountAccountQueryAPIResponse)
	},
}

// GetTaobaoOnebpDkxAccountAccountQueryAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxAccountAccountQueryAPIResponse
func GetTaobaoOnebpDkxAccountAccountQueryAPIResponse() *TaobaoOnebpDkxAccountAccountQueryAPIResponse {
	return poolTaobaoOnebpDkxAccountAccountQueryAPIResponse.Get().(*TaobaoOnebpDkxAccountAccountQueryAPIResponse)
}

// ReleaseTaobaoOnebpDkxAccountAccountQueryAPIResponse 将 TaobaoOnebpDkxAccountAccountQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxAccountAccountQueryAPIResponse(v *TaobaoOnebpDkxAccountAccountQueryAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxAccountAccountQueryAPIResponse.Put(v)
}
