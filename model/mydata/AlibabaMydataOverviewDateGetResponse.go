package mydata

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
我的效果-获取数据周期 API返回值 
alibaba.mydata.overview.date.get

获取数据管家我的效果API可以使用的数据周期
*/
type AlibabaMydataOverviewDateGetAPIResponse struct {
    model.CommonResponse
    AlibabaMydataOverviewDateGetResponse
}

// 我的效果-获取数据周期 成功返回结果
type AlibabaMydataOverviewDateGetResponse struct {
    XMLName xml.Name `xml:"alibaba_mydata_overview_date_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 我的效果可选数据周期列表
    ResultList   []DateRange `json:"result_list,omitempty" xml:"result_list>date_range,omitempty"`
}
