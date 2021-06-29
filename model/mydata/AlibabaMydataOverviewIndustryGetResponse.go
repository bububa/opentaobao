package mydata

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
我的效果-获取Top行业列表 APIResponse
alibaba.mydata.overview.industry.get

获取数据管家我的效果API可以使用的行业
*/
type AlibabaMydataOverviewIndustryGetAPIResponse struct {
    model.CommonResponse
    AlibabaMydataOverviewIndustryGetResponse
}

type AlibabaMydataOverviewIndustryGetResponse struct {
    XMLName xml.Name `xml:"alibaba_mydata_overview_industry_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 供应商Top行业列表
    
    ResultList   []Industry `json:"result_list,omitempty" xml:"result_list>industry,omitempty"`
    
    
}
