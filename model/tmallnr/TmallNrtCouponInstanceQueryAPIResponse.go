package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtCouponInstanceQueryAPIResponse 查询用户券实例 API返回值
// tmall.nrt.coupon.instance.query
//
// 查询用户券实例
type TmallNrtCouponInstanceQueryAPIResponse struct {
	model.CommonResponse
	TmallNrtCouponInstanceQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtCouponInstanceQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtCouponInstanceQueryAPIResponseModel).Reset()
}

// TmallNrtCouponInstanceQueryAPIResponseModel is 查询用户券实例 成功返回结果
type TmallNrtCouponInstanceQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_coupon_instance_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallNrtCouponInstanceQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtCouponInstanceQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrtCouponInstanceQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtCouponInstanceQueryAPIResponse)
	},
}

// GetTmallNrtCouponInstanceQueryAPIResponse 从 sync.Pool 获取 TmallNrtCouponInstanceQueryAPIResponse
func GetTmallNrtCouponInstanceQueryAPIResponse() *TmallNrtCouponInstanceQueryAPIResponse {
	return poolTmallNrtCouponInstanceQueryAPIResponse.Get().(*TmallNrtCouponInstanceQueryAPIResponse)
}

// ReleaseTmallNrtCouponInstanceQueryAPIResponse 将 TmallNrtCouponInstanceQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtCouponInstanceQueryAPIResponse(v *TmallNrtCouponInstanceQueryAPIResponse) {
	v.Reset()
	poolTmallNrtCouponInstanceQueryAPIResponse.Put(v)
}
