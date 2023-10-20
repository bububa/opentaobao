package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocloudprintisvtemplatesgetAPIRequest 获取商家使用的标准模板 API请求
// cainiao.cloudprint.isvtemplates.get
//
// 获取商家使用的标准模板
type CainiaocloudprintisvtemplatesgetAPIRequest struct {
	model.Params
}

// NewCainiaocloudprintisvtemplatesgetRequest 初始化CainiaocloudprintisvtemplatesgetAPIRequest对象
func NewCainiaocloudprintisvtemplatesgetRequest() *CainiaocloudprintisvtemplatesgetAPIRequest {
	return &CainiaocloudprintisvtemplatesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocloudprintisvtemplatesgetAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.isvtemplates.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocloudprintisvtemplatesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocloudprintisvtemplatesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
