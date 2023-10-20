package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBmCouponQueryAPIResponse 淘鲜达券信息查询接口 API返回值
// alibaba.wdk.bm.coupon.query
//
// 淘鲜达品牌营销的券信息查询接口，基于券id查询券相关信息：券id、券名称、分摊信息、面额、创建时间、开始时间、结束时间
type AlibabaWdkBmCouponQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkBmCouponQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkBmCouponQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkBmCouponQueryAPIResponseModel).Reset()
}

// AlibabaWdkBmCouponQueryAPIResponseModel is 淘鲜达券信息查询接口 成功返回结果
type AlibabaWdkBmCouponQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_bm_coupon_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BmResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkBmCouponQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkBmCouponQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkBmCouponQueryAPIResponse)
	},
}

// GetAlibabaWdkBmCouponQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkBmCouponQueryAPIResponse
func GetAlibabaWdkBmCouponQueryAPIResponse() *AlibabaWdkBmCouponQueryAPIResponse {
	return poolAlibabaWdkBmCouponQueryAPIResponse.Get().(*AlibabaWdkBmCouponQueryAPIResponse)
}

// ReleaseAlibabaWdkBmCouponQueryAPIResponse 将 AlibabaWdkBmCouponQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkBmCouponQueryAPIResponse(v *AlibabaWdkBmCouponQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkBmCouponQueryAPIResponse.Put(v)
}
