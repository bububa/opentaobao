package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcAuthGetAPIResponse TMC授权token API返回值
// taobao.tmc.auth.get
//
// TMC连接授权Token
type TaobaoTmcAuthGetAPIResponse struct {
	model.CommonResponse
	TaobaoTmcAuthGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTmcAuthGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTmcAuthGetAPIResponseModel).Reset()
}

// TaobaoTmcAuthGetAPIResponseModel is TMC授权token 成功返回结果
type TaobaoTmcAuthGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_auth_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTmcAuthGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoTmcAuthGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTmcAuthGetAPIResponse)
	},
}

// GetTaobaoTmcAuthGetAPIResponse 从 sync.Pool 获取 TaobaoTmcAuthGetAPIResponse
func GetTaobaoTmcAuthGetAPIResponse() *TaobaoTmcAuthGetAPIResponse {
	return poolTaobaoTmcAuthGetAPIResponse.Get().(*TaobaoTmcAuthGetAPIResponse)
}

// ReleaseTaobaoTmcAuthGetAPIResponse 将 TaobaoTmcAuthGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTmcAuthGetAPIResponse(v *TaobaoTmcAuthGetAPIResponse) {
	v.Reset()
	poolTaobaoTmcAuthGetAPIResponse.Put(v)
}
