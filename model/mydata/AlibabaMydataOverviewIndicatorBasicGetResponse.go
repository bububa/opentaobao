package mydata

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
我的效果-获取公司询盘流量行业表现 API返回值 
alibaba.mydata.overview.indicator.basic.get

获取公司询盘流量行业表现
*/
type AlibabaMydataOverviewIndicatorBasicGetAPIResponse struct {
    model.CommonResponse
    AlibabaMydataOverviewIndicatorBasicGetResponse
}

// 我的效果-获取公司询盘流量行业表现 成功返回结果
type AlibabaMydataOverviewIndicatorBasicGetResponse struct {
    XMLName xml.Name `xml:"alibaba_mydata_overview_indicator_basic_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 公司询盘流量指标
    Result   *CompanyIndicators `json:"result,omitempty" xml:"result,omitempty"`
}
