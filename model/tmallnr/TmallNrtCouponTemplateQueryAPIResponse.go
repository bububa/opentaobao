package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtcoupontemplatequeryAPIResponse 查询券模版 API返回值
// tmall.nrt.coupon.template.query
//
// 查询券模版
type TmallnrtcoupontemplatequeryAPIResponse struct {
	model.CommonResponse
	TmallnrtcoupontemplatequeryAPIResponseModel
}

// TmallnrtcoupontemplatequeryAPIResponseModel is 查询券模版 成功返回结果
type TmallnrtcoupontemplatequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_coupon_template_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallnrtcoupontemplatequeryResult `json:"result,omitempty" xml:"result,omitempty"`
}
