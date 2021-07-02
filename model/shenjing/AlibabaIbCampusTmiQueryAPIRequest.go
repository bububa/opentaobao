package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIbCampusTmiQueryAPIRequest IB智慧园区-查询TMI流水 API请求
// alibaba.ib.campus.tmi.query
//
// 获取特定银行账户的银行流水
type AlibabaIbCampusTmiQueryAPIRequest struct {
	model.Params
	// 查询参数
	_accountQueryReqDto *AccountQueryReqDto
}

// NewAlibabaIbCampusTmiQueryRequest 初始化AlibabaIbCampusTmiQueryAPIRequest对象
func NewAlibabaIbCampusTmiQueryRequest() *AlibabaIbCampusTmiQueryAPIRequest {
	return &AlibabaIbCampusTmiQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIbCampusTmiQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ib.campus.tmi.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIbCampusTmiQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAccountQueryReqDto is AccountQueryReqDto Setter
// 查询参数
func (r *AlibabaIbCampusTmiQueryAPIRequest) SetAccountQueryReqDto(_accountQueryReqDto *AccountQueryReqDto) error {
	r._accountQueryReqDto = _accountQueryReqDto
	r.Set("account_query_req_dto", _accountQueryReqDto)
	return nil
}

// GetAccountQueryReqDto AccountQueryReqDto Getter
func (r AlibabaIbCampusTmiQueryAPIRequest) GetAccountQueryReqDto() *AccountQueryReqDto {
	return r._accountQueryReqDto
}
