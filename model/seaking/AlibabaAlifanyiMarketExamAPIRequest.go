package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlifanyiMarketExamAPIRequest 通过考试用户 API请求
// alibaba.alifanyi.market.exam
//
// 企业或组织购买软件服务后可参与阿里翻译在线系统的考试认证，接口返回该企业或组织认证通过的用户
type AlibabaAlifanyiMarketExamAPIRequest struct {
	model.Params
	// 请求参数
	_reportQueryApiDTO *ReportQueryApiDto
}

// NewAlibabaAlifanyiMarketExamRequest 初始化AlibabaAlifanyiMarketExamAPIRequest对象
func NewAlibabaAlifanyiMarketExamRequest() *AlibabaAlifanyiMarketExamAPIRequest {
	return &AlibabaAlifanyiMarketExamAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlifanyiMarketExamAPIRequest) GetApiMethodName() string {
	return "alibaba.alifanyi.market.exam"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlifanyiMarketExamAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ReportQueryApiDTO Setter
// 请求参数
func (r *AlibabaAlifanyiMarketExamAPIRequest) SetReportQueryApiDTO(_reportQueryApiDTO *ReportQueryApiDto) error {
	r._reportQueryApiDTO = _reportQueryApiDTO
	r.Set("report_query_api_d_t_o", _reportQueryApiDTO)
	return nil
}

// Get ReportQueryApiDTO Getter
func (r AlibabaAlifanyiMarketExamAPIRequest) GetReportQueryApiDTO() *ReportQueryApiDto {
	return r._reportQueryApiDTO
}
