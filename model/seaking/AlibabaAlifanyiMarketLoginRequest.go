package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
登陆用户 APIRequest
alibaba.alifanyi.market.login

企业或组织购买软件服务后可登陆阿里翻译众包系统，接口返回该企业的用户。
*/
type AlibabaAlifanyiMarketLoginRequest struct {
    model.Params

    // 请求参数
    reportQueryApiDTO   *ReportQueryApiDto 

}

func NewAlibabaAlifanyiMarketLoginRequest() *AlibabaAlifanyiMarketLoginRequest{
    return &AlibabaAlifanyiMarketLoginRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlifanyiMarketLoginRequest) GetApiMethodName() string {
    return "alibaba.alifanyi.market.login"
}

func (r AlibabaAlifanyiMarketLoginRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlifanyiMarketLoginRequest) SetReportQueryApiDTO(reportQueryApiDTO *ReportQueryApiDto) error {
    r.reportQueryApiDTO = reportQueryApiDTO
    r.Set("report_query_api_d_t_o", reportQueryApiDTO)
    return nil
}

func (r AlibabaAlifanyiMarketLoginRequest) GetReportQueryApiDTO() *ReportQueryApiDto {
    return r.reportQueryApiDTO
}

