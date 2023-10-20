package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenpushprojectAPIRequest 大麦换验平台-第三方对外开放-项目接口pushProject API请求
// alibaba.damai.mev.open.pushproject
//
// pushProject
type AlibabadamaimevopenpushprojectAPIRequest struct {
	model.Params
	// 入参pushProjectParam
	_pushProjectParam *ThirdProjectPushOpenParam
}

// NewAlibabadamaimevopenpushprojectRequest 初始化AlibabadamaimevopenpushprojectAPIRequest对象
func NewAlibabadamaimevopenpushprojectRequest() *AlibabadamaimevopenpushprojectAPIRequest {
	return &AlibabadamaimevopenpushprojectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopenpushprojectAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushproject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopenpushprojectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopenpushprojectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushProjectParam is PushProjectParam Setter
// 入参pushProjectParam
func (r *AlibabadamaimevopenpushprojectAPIRequest) SetPushProjectParam(_pushProjectParam *ThirdProjectPushOpenParam) error {
	r._pushProjectParam = _pushProjectParam
	r.Set("push_project_param", _pushProjectParam)
	return nil
}

// GetPushProjectParam PushProjectParam Getter
func (r AlibabadamaimevopenpushprojectAPIRequest) GetPushProjectParam() *ThirdProjectPushOpenParam {
	return r._pushProjectParam
}
