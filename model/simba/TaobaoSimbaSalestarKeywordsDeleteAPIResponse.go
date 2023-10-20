package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarKeywordsDeleteAPIResponse 销量明星关键词删除 API返回值
// taobao.simba.salestar.keywords.delete
//
// （新）关键词删除相关接口
type TaobaoSimbaSalestarKeywordsDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaSalestarKeywordsDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarKeywordsDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaSalestarKeywordsDeleteAPIResponseModel).Reset()
}

// TaobaoSimbaSalestarKeywordsDeleteAPIResponseModel is 销量明星关键词删除 成功返回结果
type TaobaoSimbaSalestarKeywordsDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_salestar_keywords_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功删除条数
	Results int64 `json:"results,omitempty" xml:"results,omitempty"`
	// 删除成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarKeywordsDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = 0
	m.IsSuccess = false
}

var poolTaobaoSimbaSalestarKeywordsDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaSalestarKeywordsDeleteAPIResponse)
	},
}

// GetTaobaoSimbaSalestarKeywordsDeleteAPIResponse 从 sync.Pool 获取 TaobaoSimbaSalestarKeywordsDeleteAPIResponse
func GetTaobaoSimbaSalestarKeywordsDeleteAPIResponse() *TaobaoSimbaSalestarKeywordsDeleteAPIResponse {
	return poolTaobaoSimbaSalestarKeywordsDeleteAPIResponse.Get().(*TaobaoSimbaSalestarKeywordsDeleteAPIResponse)
}

// ReleaseTaobaoSimbaSalestarKeywordsDeleteAPIResponse 将 TaobaoSimbaSalestarKeywordsDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaSalestarKeywordsDeleteAPIResponse(v *TaobaoSimbaSalestarKeywordsDeleteAPIResponse) {
	v.Reset()
	poolTaobaoSimbaSalestarKeywordsDeleteAPIResponse.Put(v)
}
