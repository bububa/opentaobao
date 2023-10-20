package mydata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMydataOverviewIndicatorBasicGetAPIResponse 我的效果-获取公司询盘流量行业表现 API返回值
// alibaba.mydata.overview.indicator.basic.get
//
// 获取公司询盘流量行业表现
type AlibabaMydataOverviewIndicatorBasicGetAPIResponse struct {
	model.CommonResponse
	AlibabaMydataOverviewIndicatorBasicGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMydataOverviewIndicatorBasicGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMydataOverviewIndicatorBasicGetAPIResponseModel).Reset()
}

// AlibabaMydataOverviewIndicatorBasicGetAPIResponseModel is 我的效果-获取公司询盘流量行业表现 成功返回结果
type AlibabaMydataOverviewIndicatorBasicGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mydata_overview_indicator_basic_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 公司询盘流量指标
	Result *CompanyIndicators `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMydataOverviewIndicatorBasicGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMydataOverviewIndicatorBasicGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMydataOverviewIndicatorBasicGetAPIResponse)
	},
}

// GetAlibabaMydataOverviewIndicatorBasicGetAPIResponse 从 sync.Pool 获取 AlibabaMydataOverviewIndicatorBasicGetAPIResponse
func GetAlibabaMydataOverviewIndicatorBasicGetAPIResponse() *AlibabaMydataOverviewIndicatorBasicGetAPIResponse {
	return poolAlibabaMydataOverviewIndicatorBasicGetAPIResponse.Get().(*AlibabaMydataOverviewIndicatorBasicGetAPIResponse)
}

// ReleaseAlibabaMydataOverviewIndicatorBasicGetAPIResponse 将 AlibabaMydataOverviewIndicatorBasicGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMydataOverviewIndicatorBasicGetAPIResponse(v *AlibabaMydataOverviewIndicatorBasicGetAPIResponse) {
	v.Reset()
	poolAlibabaMydataOverviewIndicatorBasicGetAPIResponse.Put(v)
}
