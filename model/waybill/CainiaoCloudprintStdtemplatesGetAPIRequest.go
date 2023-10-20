package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocloudprintstdtemplatesgetAPIRequest 获取所有的菜鸟标准电子面单模板 API请求
// cainiao.cloudprint.stdtemplates.get
//
// 获取菜鸟标准电子面单模板
type CainiaocloudprintstdtemplatesgetAPIRequest struct {
	model.Params
}

// NewCainiaocloudprintstdtemplatesgetRequest 初始化CainiaocloudprintstdtemplatesgetAPIRequest对象
func NewCainiaocloudprintstdtemplatesgetRequest() *CainiaocloudprintstdtemplatesgetAPIRequest {
	return &CainiaocloudprintstdtemplatesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocloudprintstdtemplatesgetAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.stdtemplates.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocloudprintstdtemplatesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocloudprintstdtemplatesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
