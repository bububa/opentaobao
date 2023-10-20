package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse 建议默认关键词包 API返回值
// taobao.universalbp.wordpackage.suggestdefaultlist
//
// 入参推广信息，出参建议的默认关键词包
type TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponseModel).Reset()
}

// TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponseModel is 建议默认关键词包 成功返回结果
type TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_wordpackage_suggestdefaultlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpWordpackageSuggestdefaultlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse)
	},
}

// GetTaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse
func GetTaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse() *TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse {
	return poolTaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse.Get().(*TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse)
}

// ReleaseTaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse 将 TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse(v *TaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpWordpackageSuggestdefaultlistAPIResponse.Put(v)
}
