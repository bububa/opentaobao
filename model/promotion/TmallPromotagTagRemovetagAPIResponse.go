package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotagTagRemovetagAPIResponse 删除标签定义 API返回值
// tmall.promotag.tag.removetag
//
// 用于删除标签定义，但是要确保目前该标签没有人群在使用。
type TmallPromotagTagRemovetagAPIResponse struct {
	model.CommonResponse
	TmallPromotagTagRemovetagAPIResponseModel
}

// Reset 清空结构体
func (m *TmallPromotagTagRemovetagAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallPromotagTagRemovetagAPIResponseModel).Reset()
}

// TmallPromotagTagRemovetagAPIResponseModel is 删除标签定义 成功返回结果
type TmallPromotagTagRemovetagAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_promotag_tag_removetag_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallPromotagTagRemovetagAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTmallPromotagTagRemovetagAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallPromotagTagRemovetagAPIResponse)
	},
}

// GetTmallPromotagTagRemovetagAPIResponse 从 sync.Pool 获取 TmallPromotagTagRemovetagAPIResponse
func GetTmallPromotagTagRemovetagAPIResponse() *TmallPromotagTagRemovetagAPIResponse {
	return poolTmallPromotagTagRemovetagAPIResponse.Get().(*TmallPromotagTagRemovetagAPIResponse)
}

// ReleaseTmallPromotagTagRemovetagAPIResponse 将 TmallPromotagTagRemovetagAPIResponse 保存到 sync.Pool
func ReleaseTmallPromotagTagRemovetagAPIResponse(v *TmallPromotagTagRemovetagAPIResponse) {
	v.Reset()
	poolTmallPromotagTagRemovetagAPIResponse.Put(v)
}
