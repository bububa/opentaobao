package traderate

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTraderateItemtagsGetAPIResponse 通过商品ID获取标签列表 API返回值
// tmall.traderate.itemtags.get
//
// 通过商品ID获取标签详细信息
type TmallTraderateItemtagsGetAPIResponse struct {
	model.CommonResponse
	TmallTraderateItemtagsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTraderateItemtagsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTraderateItemtagsGetAPIResponseModel).Reset()
}

// TmallTraderateItemtagsGetAPIResponseModel is 通过商品ID获取标签列表 成功返回结果
type TmallTraderateItemtagsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_traderate_itemtags_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 标签列表
	Tags []TmallRateTagDetail `json:"tags,omitempty" xml:"tags>tmall_rate_tag_detail,omitempty"`
}

// Reset 清空结构体
func (m *TmallTraderateItemtagsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Tags = m.Tags[:0]
}

var poolTmallTraderateItemtagsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTraderateItemtagsGetAPIResponse)
	},
}

// GetTmallTraderateItemtagsGetAPIResponse 从 sync.Pool 获取 TmallTraderateItemtagsGetAPIResponse
func GetTmallTraderateItemtagsGetAPIResponse() *TmallTraderateItemtagsGetAPIResponse {
	return poolTmallTraderateItemtagsGetAPIResponse.Get().(*TmallTraderateItemtagsGetAPIResponse)
}

// ReleaseTmallTraderateItemtagsGetAPIResponse 将 TmallTraderateItemtagsGetAPIResponse 保存到 sync.Pool
func ReleaseTmallTraderateItemtagsGetAPIResponse(v *TmallTraderateItemtagsGetAPIResponse) {
	v.Reset()
	poolTmallTraderateItemtagsGetAPIResponse.Put(v)
}
