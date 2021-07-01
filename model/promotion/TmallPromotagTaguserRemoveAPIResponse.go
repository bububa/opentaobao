package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallPromotagTaguserRemoveAPIResponse
给用户移除优惠标签 API返回值
tmall.promotag.taguser.remove

给用户载体去标 */
type TmallPromotagTaguserRemoveAPIResponse struct {
	model.CommonResponse
	TmallPromotagTaguserRemoveAPIResponseModel
}

// TmallPromotagTaguserRemoveAPIResponseModel is 给用户移除优惠标签 成功返回结果
type TmallPromotagTaguserRemoveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_promotag_taguser_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 打标结果是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
