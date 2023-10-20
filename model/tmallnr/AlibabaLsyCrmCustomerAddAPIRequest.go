package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmcustomeraddAPIRequest 私域导购添加活动留资入口 API请求
// alibaba.lsy.crm.customer.add
//
// 私域导购添加活动留资入口
type AlibabalsycrmcustomeraddAPIRequest struct {
	model.Params
	// 入参对象
	_reqDto *NrtCrmCreateCustomerReq
}

// NewAlibabalsycrmcustomeraddRequest 初始化AlibabalsycrmcustomeraddAPIRequest对象
func NewAlibabalsycrmcustomeraddRequest() *AlibabalsycrmcustomeraddAPIRequest {
	return &AlibabalsycrmcustomeraddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsycrmcustomeraddAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.customer.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsycrmcustomeraddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsycrmcustomeraddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReqDto is ReqDto Setter
// 入参对象
func (r *AlibabalsycrmcustomeraddAPIRequest) SetReqDto(_reqDto *NrtCrmCreateCustomerReq) error {
	r._reqDto = _reqDto
	r.Set("req_dto", _reqDto)
	return nil
}

// GetReqDto ReqDto Getter
func (r AlibabalsycrmcustomeraddAPIRequest) GetReqDto() *NrtCrmCreateCustomerReq {
	return r._reqDto
}
