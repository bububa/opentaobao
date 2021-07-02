package promotion

import (
	"encoding/xml"

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

// TmallPromotagTagApplyAPIResponseModel is 优惠标签申请 成功返回结果
type TmallPromotagTagApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_promotag_tag_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否设置成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 优惠标签ID
	TagId int64 `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
}
