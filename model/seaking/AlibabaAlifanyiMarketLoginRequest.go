package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
登陆用户 API请求
alibaba.alifanyi.market.login

企业或组织购买软件服务后可登陆阿里翻译众包系统，接口返回该企业的用户。
*/
type AlibabaAlifanyiMarketLoginRequest struct {
    model.Params
    // 请求参数
    _reportQueryApiDTO   *ReportQueryApiDto
}

// 初始化AlibabaAlifanyiMarketLoginRequest对象
func NewAlibabaAlifanyiMarketLoginRequest() *AlibabaAlifanyiMarketLoginRequest{
    return &AlibabaAlifanyiMarketLoginRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlifanyiMarketLoginRequest) GetApiMethodName() string {
    return "alibaba.alifanyi.market.login"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlifanyiMarketLoginRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReportQueryApiDTO Setter
// 请求参数
func (r *AlibabaAlifanyiMarketLoginRequest) SetReportQueryApiDTO(_reportQueryApiDTO *ReportQueryApiDto) error {
    r._reportQueryApiDTO = _reportQueryApiDTO
    r.Set("report_query_api_d_t_o", _reportQueryApiDTO)
    return nil
}

// ReportQueryApiDTO Getter
func (r AlibabaAlifanyiMarketLoginRequest) GetReportQueryApiDTO() *ReportQueryApiDto {
    return r._reportQueryApiDTO
}