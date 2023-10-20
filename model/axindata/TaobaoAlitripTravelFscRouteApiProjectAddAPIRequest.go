package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiprojectaddAPIRequest 新增团期 API请求
// taobao.alitrip.travel.fsc.route.api.project.add
//
// 新增团期
type TaobaoalitriptravelfscrouteapiprojectaddAPIRequest struct {
	model.Params
	// fscProjectModifyRequest
	_fscProjectModifyRequest *FscProjectModifyRequest
}

// NewTaobaoalitriptravelfscrouteapiprojectaddRequest 初始化TaobaoalitriptravelfscrouteapiprojectaddAPIRequest对象
func NewTaobaoalitriptravelfscrouteapiprojectaddRequest() *TaobaoalitriptravelfscrouteapiprojectaddAPIRequest {
	return &TaobaoalitriptravelfscrouteapiprojectaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfscrouteapiprojectaddAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.project.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfscrouteapiprojectaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfscrouteapiprojectaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProjectModifyRequest is FscProjectModifyRequest Setter
// fscProjectModifyRequest
func (r *TaobaoalitriptravelfscrouteapiprojectaddAPIRequest) SetFscProjectModifyRequest(_fscProjectModifyRequest *FscProjectModifyRequest) error {
	r._fscProjectModifyRequest = _fscProjectModifyRequest
	r.Set("fsc_project_modify_request", _fscProjectModifyRequest)
	return nil
}

// GetFscProjectModifyRequest FscProjectModifyRequest Getter
func (r TaobaoalitriptravelfscrouteapiprojectaddAPIRequest) GetFscProjectModifyRequest() *FscProjectModifyRequest {
	return r._fscProjectModifyRequest
}
