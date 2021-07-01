package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenPushprojectAPIRequest
大麦换验平台-第三方对外开放-项目接口pushProject API请求
alibaba.damai.mev.open.pushproject

pushProject */
type AlibabaDamaiMevOpenPushprojectAPIRequest struct {
	model.Params
	// 入参pushProjectParam
	_pushProjectParam *ThirdProjectPushOpenParam
}

// NewAlibabaDamaiMevOpenPushprojectRequest 初始化AlibabaDamaiMevOpenPushprojectAPIRequest对象
func NewAlibabaDamaiMevOpenPushprojectRequest() *AlibabaDamaiMevOpenPushprojectAPIRequest {
	return &AlibabaDamaiMevOpenPushprojectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushprojectAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushproject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushprojectAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PushProjectParam Setter
// 入参pushProjectParam
func (r *AlibabaDamaiMevOpenPushprojectAPIRequest) SetPushProjectParam(_pushProjectParam *ThirdProjectPushOpenParam) error {
	r._pushProjectParam = _pushProjectParam
	r.Set("push_project_param", _pushProjectParam)
	return nil
}

// Get PushProjectParam Getter
func (r AlibabaDamaiMevOpenPushprojectAPIRequest) GetPushProjectParam() *ThirdProjectPushOpenParam {
	return r._pushProjectParam
}
