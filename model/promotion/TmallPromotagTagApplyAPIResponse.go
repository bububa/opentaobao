package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotagTagApplyAPIResponse 优惠标签申请 API返回值
// tmall.promotag.tag.apply
//
// 创建优惠标签
type TmallPromotagTagApplyAPIResponse struct {
	model.CommonResponse
	TmallPromotagTagApplyAPIResponseModel
}

// Reset 清空结构体
func (m *TmallPromotagTagApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallPromotagTagApplyAPIResponseModel).Reset()
}

// TmallPromotagTagApplyAPIResponseModel is 优惠标签申请 成功返回结果
type TmallPromotagTagApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_promotag_tag_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠标签ID
	TagId int64 `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
	// 是否设置成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallPromotagTagApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TagId = 0
	m.IsSuccess = false
}

var poolTmallPromotagTagApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallPromotagTagApplyAPIResponse)
	},
}

// GetTmallPromotagTagApplyAPIResponse 从 sync.Pool 获取 TmallPromotagTagApplyAPIResponse
func GetTmallPromotagTagApplyAPIResponse() *TmallPromotagTagApplyAPIResponse {
	return poolTmallPromotagTagApplyAPIResponse.Get().(*TmallPromotagTagApplyAPIResponse)
}

// ReleaseTmallPromotagTagApplyAPIResponse 将 TmallPromotagTagApplyAPIResponse 保存到 sync.Pool
func ReleaseTmallPromotagTagApplyAPIResponse(v *TmallPromotagTagApplyAPIResponse) {
	v.Reset()
	poolTmallPromotagTagApplyAPIResponse.Put(v)
}
