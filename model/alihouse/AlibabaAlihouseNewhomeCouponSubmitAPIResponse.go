package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeCouponSubmitAPIResponse 提交专车优惠券活动 API返回值
// alibaba.alihouse.newhome.coupon.submit
//
// 提交专车优惠券活动
type AlibabaAlihouseNewhomeCouponSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeCouponSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeCouponSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeCouponSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeCouponSubmitAPIResponseModel is 提交专车优惠券活动 成功返回结果
type AlibabaAlihouseNewhomeCouponSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_coupon_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeCouponSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeCouponSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeCouponSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeCouponSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeCouponSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeCouponSubmitAPIResponse
func GetAlibabaAlihouseNewhomeCouponSubmitAPIResponse() *AlibabaAlihouseNewhomeCouponSubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeCouponSubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeCouponSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeCouponSubmitAPIResponse 将 AlibabaAlihouseNewhomeCouponSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeCouponSubmitAPIResponse(v *AlibabaAlihouseNewhomeCouponSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeCouponSubmitAPIResponse.Put(v)
}
