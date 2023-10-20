package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappdistributionmaterialgetAPIRequest 小程序投放---读取投放入口素材信息 API请求
// taobao.miniapp.distribution.material.get
//
// 读取已录入的投放入口素材信息。
type TaobaominiappdistributionmaterialgetAPIRequest struct {
	model.Params
	// 投放入口素材信息
	_materialRequest *MiniAppEntranceMaterialBizOpenDto
}

// NewTaobaominiappdistributionmaterialgetRequest 初始化TaobaominiappdistributionmaterialgetAPIRequest对象
func NewTaobaominiappdistributionmaterialgetRequest() *TaobaominiappdistributionmaterialgetAPIRequest {
	return &TaobaominiappdistributionmaterialgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappdistributionmaterialgetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.material.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappdistributionmaterialgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappdistributionmaterialgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMaterialRequest is MaterialRequest Setter
// 投放入口素材信息
func (r *TaobaominiappdistributionmaterialgetAPIRequest) SetMaterialRequest(_materialRequest *MiniAppEntranceMaterialBizOpenDto) error {
	r._materialRequest = _materialRequest
	r.Set("material_request", _materialRequest)
	return nil
}

// GetMaterialRequest MaterialRequest Getter
func (r TaobaominiappdistributionmaterialgetAPIRequest) GetMaterialRequest() *MiniAppEntranceMaterialBizOpenDto {
	return r._materialRequest
}
