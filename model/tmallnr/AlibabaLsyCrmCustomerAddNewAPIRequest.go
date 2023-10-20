package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmcustomeraddnewAPIRequest 导购域留资接口 API请求
// alibaba.lsy.crm.customer.add.new
//
// 导购域提供留资入口
type AlibabalsycrmcustomeraddnewAPIRequest struct {
	model.Params
	// 入参DTO
	_reqDTO *NrtCrmCreateCustomerReq
}

// NewAlibabalsycrmcustomeraddnewRequest 初始化AlibabalsycrmcustomeraddnewAPIRequest对象
func NewAlibabalsycrmcustomeraddnewRequest() *AlibabalsycrmcustomeraddnewAPIRequest {
	return &AlibabalsycrmcustomeraddnewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsycrmcustomeraddnewAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.customer.add.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsycrmcustomeraddnewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsycrmcustomeraddnewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReqDTO is ReqDTO Setter
// 入参DTO
func (r *AlibabalsycrmcustomeraddnewAPIRequest) SetReqDTO(_reqDTO *NrtCrmCreateCustomerReq) error {
	r._reqDTO = _reqDTO
	r.Set("req_d_t_o", _reqDTO)
	return nil
}

// GetReqDTO ReqDTO Getter
func (r AlibabalsycrmcustomeraddnewAPIRequest) GetReqDTO() *NrtCrmCreateCustomerReq {
	return r._reqDTO
}
