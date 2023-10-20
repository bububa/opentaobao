package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkumsordergetAPIRequest 查询店仓作业单据清单 （库存对账辅助）-回流单 API请求
// alibaba.wdk.ums.order.get
//
// 查询店仓作业单据清单 （库存对账辅助）-回流单
type AlibabawdkumsordergetAPIRequest struct {
	model.Params
	// 查询单据的dto
	_queryErpbillDto *QueryErpBillDto
}

// NewAlibabawdkumsordergetRequest 初始化AlibabawdkumsordergetAPIRequest对象
func NewAlibabawdkumsordergetRequest() *AlibabawdkumsordergetAPIRequest {
	return &AlibabawdkumsordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkumsordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkumsordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkumsordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryErpbillDto is QueryErpbillDto Setter
// 查询单据的dto
func (r *AlibabawdkumsordergetAPIRequest) SetQueryErpbillDto(_queryErpbillDto *QueryErpBillDto) error {
	r._queryErpbillDto = _queryErpbillDto
	r.Set("query_erpbill_dto", _queryErpbillDto)
	return nil
}

// GetQueryErpbillDto QueryErpbillDto Getter
func (r AlibabawdkumsordergetAPIRequest) GetQueryErpbillDto() *QueryErpBillDto {
	return r._queryErpbillDto
}
