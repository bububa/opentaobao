package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过考试用户 APIRequest
alibaba.alifanyi.market.exam

企业或组织购买软件服务后可参与阿里翻译在线系统的考试认证，接口返回该企业或组织认证通过的用户
*/
type AlibabaAlifanyiMarketExamRequest struct {
    model.Params

    // 请求参数
    reportQueryApiDTO   *ReportQueryApiDto 

}

func NewAlibabaAlifanyiMarketExamRequest() *AlibabaAlifanyiMarketExamRequest{
    return &AlibabaAlifanyiMarketExamRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlifanyiMarketExamRequest) GetApiMethodName() string {
    return "alibaba.alifanyi.market.exam"
}

func (r AlibabaAlifanyiMarketExamRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlifanyiMarketExamRequest) SetReportQueryApiDTO(reportQueryApiDTO *ReportQueryApiDto) error {
    r.reportQueryApiDTO = reportQueryApiDTO
    r.Set("report_query_api_d_t_o", reportQueryApiDTO)
    return nil
}

func (r AlibabaAlifanyiMarketExamRequest) GetReportQueryApiDTO() *ReportQueryApiDto {
    return r.reportQueryApiDTO
}

