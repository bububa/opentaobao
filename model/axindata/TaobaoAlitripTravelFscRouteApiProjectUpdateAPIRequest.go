package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiprojectupdateAPIRequest 更新团期 API请求
// taobao.alitrip.travel.fsc.route.api.project.update
//
// 更新团期
type TaobaoalitriptravelfscrouteapiprojectupdateAPIRequest struct {
	model.Params
	// fscProjectUpdateRequest
	_fscProjectUpdateRequest *FscProjectModifyRequest
}

// NewTaobaoalitriptravelfscrouteapiprojectupdateRequest 初始化TaobaoalitriptravelfscrouteapiprojectupdateAPIRequest对象
func NewTaobaoalitriptravelfscrouteapiprojectupdateRequest() *TaobaoalitriptravelfscrouteapiprojectupdateAPIRequest {
	return &TaobaoalitriptravelfscrouteapiprojectupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfscrouteapiprojectupdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.project.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfscrouteapiprojectupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfscrouteapiprojectupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProjectUpdateRequest is FscProjectUpdateRequest Setter
// fscProjectUpdateRequest
func (r *TaobaoalitriptravelfscrouteapiprojectupdateAPIRequest) SetFscProjectUpdateRequest(_fscProjectUpdateRequest *FscProjectModifyRequest) error {
	r._fscProjectUpdateRequest = _fscProjectUpdateRequest
	r.Set("fsc_project_update_request", _fscProjectUpdateRequest)
	return nil
}

// GetFscProjectUpdateRequest FscProjectUpdateRequest Getter
func (r TaobaoalitriptravelfscrouteapiprojectupdateAPIRequest) GetFscProjectUpdateRequest() *FscProjectModifyRequest {
	return r._fscProjectUpdateRequest
}
