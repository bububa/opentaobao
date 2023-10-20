package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampuscorecompanycampusgetcombycamidAPIRequest 根据园区ID获取运营公司信息 API请求
// alibaba.campus.core.companycampus.getcombycamid
//
// 根据园区ID获取运营公司信息
type AlibabacampuscorecompanycampusgetcombycamidAPIRequest struct {
	model.Params
	// WorkBenchContext
	_param0 *WorkBenchContext
}

// NewAlibabacampuscorecompanycampusgetcombycamidRequest 初始化AlibabacampuscorecompanycampusgetcombycamidAPIRequest对象
func NewAlibabacampuscorecompanycampusgetcombycamidRequest() *AlibabacampuscorecompanycampusgetcombycamidAPIRequest {
	return &AlibabacampuscorecompanycampusgetcombycamidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampuscorecompanycampusgetcombycamidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.core.companycampus.getcombycamid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampuscorecompanycampusgetcombycamidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampuscorecompanycampusgetcombycamidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// WorkBenchContext
func (r *AlibabacampuscorecompanycampusgetcombycamidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampuscorecompanycampusgetcombycamidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}
