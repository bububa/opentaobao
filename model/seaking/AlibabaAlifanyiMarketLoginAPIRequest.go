package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlifanyiMarketLoginAPIRequest 登陆用户 API请求
// alibaba.alifanyi.market.login
//
// 企业或组织购买软件服务后可登陆阿里翻译众包系统，接口返回该企业的用户。
type AlibabaAlifanyiMarketLoginAPIRequest struct {
	model.Params
	// 请求参数
	_reportQueryApiDTO *ReportQueryApiDto
}

// NewAlibabaAlifanyiMarketLoginRequest 初始化AlibabaAlifanyiMarketLoginAPIRequest对象
func NewAlibabaAlifanyiMarketLoginRequest() *AlibabaAlifanyiMarketLoginAPIRequest {
	return &AlibabaAlifanyiMarketLoginAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlifanyiMarketLoginAPIRequest) GetApiMethodName() string {
	return "alibaba.alifanyi.market.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlifanyiMarketLoginAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReportQueryApiDTO is ReportQueryApiDTO Setter
// 请求参数
func (r *AlibabaAlifanyiMarketLoginAPIRequest) SetReportQueryApiDTO(_reportQueryApiDTO *ReportQueryApiDto) error {
	r._reportQueryApiDTO = _reportQueryApiDTO
	r.Set("report_query_api_d_t_o", _reportQueryApiDTO)
	return nil
}

// GetReportQueryApiDTO ReportQueryApiDTO Getter
func (r AlibabaAlifanyiMarketLoginAPIRequest) GetReportQueryApiDTO() *ReportQueryApiDto {
	return r._reportQueryApiDTO
}
