package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocloudprintmystdtemplatesgetAPIRequest 获取用户使用的菜鸟电子面单模板信息 API请求
// cainiao.cloudprint.mystdtemplates.get
//
// 获取用户使用的菜鸟电子面单
type CainiaocloudprintmystdtemplatesgetAPIRequest struct {
	model.Params
}

// NewCainiaocloudprintmystdtemplatesgetRequest 初始化CainiaocloudprintmystdtemplatesgetAPIRequest对象
func NewCainiaocloudprintmystdtemplatesgetRequest() *CainiaocloudprintmystdtemplatesgetAPIRequest {
	return &CainiaocloudprintmystdtemplatesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocloudprintmystdtemplatesgetAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.mystdtemplates.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocloudprintmystdtemplatesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocloudprintmystdtemplatesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
