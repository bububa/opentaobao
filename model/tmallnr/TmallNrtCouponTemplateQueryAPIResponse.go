package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtCouponTemplateQueryAPIResponse 查询券模版 API返回值
// tmall.nrt.coupon.template.query
//
// 查询券模版
type TmallNrtCouponTemplateQueryAPIResponse struct {
	model.CommonResponse
	TmallNrtCouponTemplateQueryAPIResponseModel
}

// TmallNrtCouponTemplateQueryAPIResponseModel is 查询券模版 成功返回结果
type TmallNrtCouponTemplateQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_coupon_template_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallNrtCouponTemplateQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
