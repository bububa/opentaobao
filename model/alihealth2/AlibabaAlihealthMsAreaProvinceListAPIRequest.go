package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmsareaprovincelistAPIRequest 疫苗预约省份列表查询 API请求
// alibaba.alihealth.ms.area.province.list
//
// 微信小程序疫苗预约省份列表查询
type AlibabaalihealthmsareaprovincelistAPIRequest struct {
	model.Params
}

// NewAlibabaalihealthmsareaprovincelistRequest 初始化AlibabaalihealthmsareaprovincelistAPIRequest对象
func NewAlibabaalihealthmsareaprovincelistRequest() *AlibabaalihealthmsareaprovincelistAPIRequest {
	return &AlibabaalihealthmsareaprovincelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmsareaprovincelistAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.ms.area.province.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmsareaprovincelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmsareaprovincelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}
