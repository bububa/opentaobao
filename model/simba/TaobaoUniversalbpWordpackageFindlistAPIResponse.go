package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpWordpackageFindlistAPIResponse 词包列表查询 API返回值
// taobao.universalbp.wordpackage.findlist
//
// 根据计划+单元id，查绑定的词包列表
type TaobaoUniversalbpWordpackageFindlistAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpWordpackageFindlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpWordpackageFindlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpWordpackageFindlistAPIResponseModel).Reset()
}

// TaobaoUniversalbpWordpackageFindlistAPIResponseModel is 词包列表查询 成功返回结果
type TaobaoUniversalbpWordpackageFindlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_wordpackage_findlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpWordpackageFindlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpWordpackageFindlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpWordpackageFindlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpWordpackageFindlistAPIResponse)
	},
}

// GetTaobaoUniversalbpWordpackageFindlistAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpWordpackageFindlistAPIResponse
func GetTaobaoUniversalbpWordpackageFindlistAPIResponse() *TaobaoUniversalbpWordpackageFindlistAPIResponse {
	return poolTaobaoUniversalbpWordpackageFindlistAPIResponse.Get().(*TaobaoUniversalbpWordpackageFindlistAPIResponse)
}

// ReleaseTaobaoUniversalbpWordpackageFindlistAPIResponse 将 TaobaoUniversalbpWordpackageFindlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpWordpackageFindlistAPIResponse(v *TaobaoUniversalbpWordpackageFindlistAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpWordpackageFindlistAPIResponse.Put(v)
}
