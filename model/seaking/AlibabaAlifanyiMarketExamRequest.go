package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过考试用户 API请求
alibaba.alifanyi.market.exam

企业或组织购买软件服务后可参与阿里翻译在线系统的考试认证，接口返回该企业或组织认证通过的用户
*/
type AlibabaAlifanyiMarketExamRequest struct {
    model.Params
    // 请求参数
    _reportQueryApiDTO   *ReportQueryApiDto
}

// 初始化AlibabaAlifanyiMarketExamRequest对象
func NewAlibabaAlifanyiMarketExamRequest() *AlibabaAlifanyiMarketExamRequest{
    return &AlibabaAlifanyiMarketExamRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlifanyiMarketExamRequest) GetApiMethodName() string {
    return "alibaba.alifanyi.market.exam"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlifanyiMarketExamRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReportQueryApiDTO Setter
// 请求参数
func (r *AlibabaAlifanyiMarketExamRequest) SetReportQueryApiDTO(_reportQueryApiDTO *ReportQueryApiDto) error {
    r._reportQueryApiDTO = _reportQueryApiDTO
    r.Set("report_query_api_d_t_o", _reportQueryApiDTO)
    return nil
}

// ReportQueryApiDTO Getter
func (r AlibabaAlifanyiMarketExamRequest) GetReportQueryApiDTO() *ReportQueryApiDto {
    return r._reportQueryApiDTO
}
