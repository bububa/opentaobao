package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVasSubscSearchAPIResponse 订购记录导出 API返回值
// taobao.vas.subsc.search
//
// 用于ISV查询自己名下的应用及收费项目的订购记录
type TaobaoVasSubscSearchAPIResponse struct {
	model.CommonResponse
	TaobaoVasSubscSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVasSubscSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVasSubscSearchAPIResponseModel).Reset()
}

// TaobaoVasSubscSearchAPIResponseModel is 订购记录导出 成功返回结果
type TaobaoVasSubscSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"vas_subsc_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订购关系对象
	ArticleSubs []ArticleSub `json:"article_subs,omitempty" xml:"article_subs>article_sub,omitempty"`
	// 总记录数
	TotalItem int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVasSubscSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ArticleSubs = m.ArticleSubs[:0]
	m.TotalItem = 0
}

var poolTaobaoVasSubscSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVasSubscSearchAPIResponse)
	},
}

// GetTaobaoVasSubscSearchAPIResponse 从 sync.Pool 获取 TaobaoVasSubscSearchAPIResponse
func GetTaobaoVasSubscSearchAPIResponse() *TaobaoVasSubscSearchAPIResponse {
	return poolTaobaoVasSubscSearchAPIResponse.Get().(*TaobaoVasSubscSearchAPIResponse)
}

// ReleaseTaobaoVasSubscSearchAPIResponse 将 TaobaoVasSubscSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVasSubscSearchAPIResponse(v *TaobaoVasSubscSearchAPIResponse) {
	v.Reset()
	poolTaobaoVasSubscSearchAPIResponse.Put(v)
}
