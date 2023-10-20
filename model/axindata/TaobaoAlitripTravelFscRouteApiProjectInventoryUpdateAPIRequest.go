package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIRequest 更新团期库存 API请求
// taobao.alitrip.travel.fsc.route.api.project.inventory.update
//
// 更新团期库存
type TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIRequest struct {
	model.Params
	// fscProjectInventoryUpdateRequest
	_fscProjectInventoryUpdateRequest *FscProjectInventoryUpdateRequest
}

// NewTaobaoalitriptravelfscrouteapiprojectinventoryupdateRequest 初始化TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIRequest对象
func NewTaobaoalitriptravelfscrouteapiprojectinventoryupdateRequest() *TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIRequest {
	return &TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.route.api.project.inventory.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFscProjectInventoryUpdateRequest is FscProjectInventoryUpdateRequest Setter
// fscProjectInventoryUpdateRequest
func (r *TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIRequest) SetFscProjectInventoryUpdateRequest(_fscProjectInventoryUpdateRequest *FscProjectInventoryUpdateRequest) error {
	r._fscProjectInventoryUpdateRequest = _fscProjectInventoryUpdateRequest
	r.Set("fsc_project_inventory_update_request", _fscProjectInventoryUpdateRequest)
	return nil
}

// GetFscProjectInventoryUpdateRequest FscProjectInventoryUpdateRequest Getter
func (r TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIRequest) GetFscProjectInventoryUpdateRequest() *FscProjectInventoryUpdateRequest {
	return r._fscProjectInventoryUpdateRequest
}
