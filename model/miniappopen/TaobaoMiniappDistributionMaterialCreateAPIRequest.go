package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappdistributionmaterialcreateAPIRequest 小程序投放--新建投放素材 API请求
// taobao.miniapp.distribution.material.create
//
// 为可投放的小程序，增加入口的素材信息，比如图片、引导文案等等。
type TaobaominiappdistributionmaterialcreateAPIRequest struct {
	model.Params
	// 投放入口素材信息
	_materialRequest *MiniAppEntranceMaterialBizOpenDto
}

// NewTaobaominiappdistributionmaterialcreateRequest 初始化TaobaominiappdistributionmaterialcreateAPIRequest对象
func NewTaobaominiappdistributionmaterialcreateRequest() *TaobaominiappdistributionmaterialcreateAPIRequest {
	return &TaobaominiappdistributionmaterialcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappdistributionmaterialcreateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.material.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappdistributionmaterialcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappdistributionmaterialcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMaterialRequest is MaterialRequest Setter
// 投放入口素材信息
func (r *TaobaominiappdistributionmaterialcreateAPIRequest) SetMaterialRequest(_materialRequest *MiniAppEntranceMaterialBizOpenDto) error {
	r._materialRequest = _materialRequest
	r.Set("material_request", _materialRequest)
	return nil
}

// GetMaterialRequest MaterialRequest Getter
func (r TaobaominiappdistributionmaterialcreateAPIRequest) GetMaterialRequest() *MiniAppEntranceMaterialBizOpenDto {
	return r._materialRequest
}
