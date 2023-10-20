package tmallnr

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TmallNrtCouponTemplateQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtCouponTemplateQueryAPIResponseModel).Reset()
}

// TmallNrtCouponTemplateQueryAPIResponseModel is 查询券模版 成功返回结果
type TmallNrtCouponTemplateQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_coupon_template_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallNrtCouponTemplateQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtCouponTemplateQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrtCouponTemplateQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtCouponTemplateQueryAPIResponse)
	},
}

// GetTmallNrtCouponTemplateQueryAPIResponse 从 sync.Pool 获取 TmallNrtCouponTemplateQueryAPIResponse
func GetTmallNrtCouponTemplateQueryAPIResponse() *TmallNrtCouponTemplateQueryAPIResponse {
	return poolTmallNrtCouponTemplateQueryAPIResponse.Get().(*TmallNrtCouponTemplateQueryAPIResponse)
}

// ReleaseTmallNrtCouponTemplateQueryAPIResponse 将 TmallNrtCouponTemplateQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtCouponTemplateQueryAPIResponse(v *TmallNrtCouponTemplateQueryAPIResponse) {
	v.Reset()
	poolTmallNrtCouponTemplateQueryAPIResponse.Put(v)
}
