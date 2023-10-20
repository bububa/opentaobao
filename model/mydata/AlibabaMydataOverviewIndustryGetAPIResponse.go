package mydata

import (
	"encoding/xml"

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

// AlibabaMydataOverviewIndustryGetAPIResponseModel is 我的效果-获取Top行业列表 成功返回结果
type AlibabaMydataOverviewIndustryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mydata_overview_industry_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 供应商Top行业列表
	ResultList []Industry `json:"result_list,omitempty" xml:"result_list>industry,omitempty"`
}
