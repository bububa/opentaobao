package mydata

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMydataOverviewIndustryGetAPIResponse 我的效果-获取Top行业列表 API返回值
// alibaba.mydata.overview.industry.get
//
// 获取数据管家我的效果API可以使用的行业
type AlibabaMydataOverviewIndustryGetAPIResponse struct {
	model.CommonResponse
	AlibabaMydataOverviewIndustryGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMydataOverviewIndustryGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMydataOverviewIndustryGetAPIResponseModel).Reset()
}

// AlibabaMydataOverviewIndustryGetAPIResponseModel is 我的效果-获取Top行业列表 成功返回结果
type AlibabaMydataOverviewIndustryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mydata_overview_industry_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 供应商Top行业列表
	ResultList []Industry `json:"result_list,omitempty" xml:"result_list>industry,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMydataOverviewIndustryGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolAlibabaMydataOverviewIndustryGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMydataOverviewIndustryGetAPIResponse)
	},
}

// GetAlibabaMydataOverviewIndustryGetAPIResponse 从 sync.Pool 获取 AlibabaMydataOverviewIndustryGetAPIResponse
func GetAlibabaMydataOverviewIndustryGetAPIResponse() *AlibabaMydataOverviewIndustryGetAPIResponse {
	return poolAlibabaMydataOverviewIndustryGetAPIResponse.Get().(*AlibabaMydataOverviewIndustryGetAPIResponse)
}

// ReleaseAlibabaMydataOverviewIndustryGetAPIResponse 将 AlibabaMydataOverviewIndustryGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMydataOverviewIndustryGetAPIResponse(v *AlibabaMydataOverviewIndustryGetAPIResponse) {
	v.Reset()
	poolAlibabaMydataOverviewIndustryGetAPIResponse.Put(v)
}
