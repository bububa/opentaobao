package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse 品牌营销导购员券推广统计数据回流 API返回值
// alibaba.txcs.brandmarketing.coupon.statistics.get
//
// 请求券统计数据回流
type AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse struct {
	model.CommonResponse
	AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponseModel).Reset()
}

// AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponseModel is 品牌营销导购员券推广统计数据回流 成功返回结果
type AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_txcs_brandmarketing_coupon_statistics_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiPageResult *ApiPageResult `json:"api_page_result,omitempty" xml:"api_page_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiPageResult = nil
}

var poolAlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse)
	},
}

// GetAlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse 从 sync.Pool 获取 AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse
func GetAlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse() *AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse {
	return poolAlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse.Get().(*AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse)
}

// ReleaseAlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse 将 AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse(v *AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse) {
	v.Reset()
	poolAlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse.Put(v)
}
