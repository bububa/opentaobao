package mydata

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
我的效果-获取数据周期 APIResponse
alibaba.mydata.overview.date.get

获取数据管家我的效果API可以使用的数据周期
*/
type AlibabaMydataOverviewDateGetAPIResponse struct {
    model.CommonResponse
    AlibabaMydataOverviewDateGetResponse
}

type AlibabaMydataOverviewDateGetResponse struct {
    XMLName xml.Name `xml:"alibaba_mydata_overview_date_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 我的效果可选数据周期列表
    
    ResultList   []DateRange `json:"result_list,omitempty" xml:"result_list>date_range,omitempty"`
    
    
}
