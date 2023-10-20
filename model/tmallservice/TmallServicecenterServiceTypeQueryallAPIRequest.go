package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterservicetypequeryallAPIRequest 服务供应链服务类型 API请求
// tmall.servicecenter.service.type.queryall
//
// 查询天猫服务类型列表
type TmallservicecenterservicetypequeryallAPIRequest struct {
	model.Params
}

// NewTmallservicecenterservicetypequeryallRequest 初始化TmallservicecenterservicetypequeryallAPIRequest对象
func NewTmallservicecenterservicetypequeryallRequest() *TmallservicecenterservicetypequeryallAPIRequest {
	return &TmallservicecenterservicetypequeryallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterservicetypequeryallAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.service.type.queryall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterservicetypequeryallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterservicetypequeryallAPIRequest) GetRawParams() model.Params {
	return r.Params
}
