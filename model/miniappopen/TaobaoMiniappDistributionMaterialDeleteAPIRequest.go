package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappdistributionmaterialdeleteAPIRequest 小程序投放 --- 删除投放素材 API请求
// taobao.miniapp.distribution.material.delete
//
// 删除已录入的投放入口素材信息。
type TaobaominiappdistributionmaterialdeleteAPIRequest struct {
	model.Params
	// 投放入口素材信息
	_materialRequest *MiniAppEntranceMaterialBizOpenDto
}

// NewTaobaominiappdistributionmaterialdeleteRequest 初始化TaobaominiappdistributionmaterialdeleteAPIRequest对象
func NewTaobaominiappdistributionmaterialdeleteRequest() *TaobaominiappdistributionmaterialdeleteAPIRequest {
	return &TaobaominiappdistributionmaterialdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappdistributionmaterialdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.material.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappdistributionmaterialdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappdistributionmaterialdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMaterialRequest is MaterialRequest Setter
// 投放入口素材信息
func (r *TaobaominiappdistributionmaterialdeleteAPIRequest) SetMaterialRequest(_materialRequest *MiniAppEntranceMaterialBizOpenDto) error {
	r._materialRequest = _materialRequest
	r.Set("material_request", _materialRequest)
	return nil
}

// GetMaterialRequest MaterialRequest Getter
func (r TaobaominiappdistributionmaterialdeleteAPIRequest) GetMaterialRequest() *MiniAppEntranceMaterialBizOpenDto {
	return r._materialRequest
}
