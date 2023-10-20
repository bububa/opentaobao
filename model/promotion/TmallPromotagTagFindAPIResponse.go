package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotagTagFindAPIResponse 查询标签接口 API返回值
// tmall.promotag.tag.find
//
// 查询用户创建的所有标签
type TmallPromotagTagFindAPIResponse struct {
	model.CommonResponse
	TmallPromotagTagFindAPIResponseModel
}

// Reset 清空结构体
func (m *TmallPromotagTagFindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallPromotagTagFindAPIResponseModel).Reset()
}

// TmallPromotagTagFindAPIResponseModel is 查询标签接口 成功返回结果
type TmallPromotagTagFindAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_promotag_tag_find_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果类型
	QueryResult *PromotionTagQuery `json:"query_result,omitempty" xml:"query_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallPromotagTagFindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.QueryResult = nil
}

var poolTmallPromotagTagFindAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallPromotagTagFindAPIResponse)
	},
}

// GetTmallPromotagTagFindAPIResponse 从 sync.Pool 获取 TmallPromotagTagFindAPIResponse
func GetTmallPromotagTagFindAPIResponse() *TmallPromotagTagFindAPIResponse {
	return poolTmallPromotagTagFindAPIResponse.Get().(*TmallPromotagTagFindAPIResponse)
}

// ReleaseTmallPromotagTagFindAPIResponse 将 TmallPromotagTagFindAPIResponse 保存到 sync.Pool
func ReleaseTmallPromotagTagFindAPIResponse(v *TmallPromotagTagFindAPIResponse) {
	v.Reset()
	poolTmallPromotagTagFindAPIResponse.Put(v)
}
