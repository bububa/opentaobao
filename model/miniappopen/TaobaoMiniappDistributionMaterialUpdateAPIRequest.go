package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappdistributionmaterialupdateAPIRequest 小程序投放 --- 更新投放素材 API请求
// taobao.miniapp.distribution.material.update
//
// 修改已录入的投放素材信息。
type TaobaominiappdistributionmaterialupdateAPIRequest struct {
	model.Params
	// 投放入口素材信息
	_materialRequest *MiniAppEntranceMaterialBizOpenDto
}

// NewTaobaominiappdistributionmaterialupdateRequest 初始化TaobaominiappdistributionmaterialupdateAPIRequest对象
func NewTaobaominiappdistributionmaterialupdateRequest() *TaobaominiappdistributionmaterialupdateAPIRequest {
	return &TaobaominiappdistributionmaterialupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappdistributionmaterialupdateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.material.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappdistributionmaterialupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappdistributionmaterialupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMaterialRequest is MaterialRequest Setter
// 投放入口素材信息
func (r *TaobaominiappdistributionmaterialupdateAPIRequest) SetMaterialRequest(_materialRequest *MiniAppEntranceMaterialBizOpenDto) error {
	r._materialRequest = _materialRequest
	r.Set("material_request", _materialRequest)
	return nil
}

// GetMaterialRequest MaterialRequest Getter
func (r TaobaominiappdistributionmaterialupdateAPIRequest) GetMaterialRequest() *MiniAppEntranceMaterialBizOpenDto {
	return r._materialRequest
}
