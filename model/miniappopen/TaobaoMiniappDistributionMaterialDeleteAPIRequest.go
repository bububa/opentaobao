package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionMaterialDeleteAPIRequest 小程序投放 --- 删除投放素材 API请求
// taobao.miniapp.distribution.material.delete
//
// 删除已录入的投放入口素材信息。
type TaobaoMiniappDistributionMaterialDeleteAPIRequest struct {
	model.Params
	// 投放入口素材信息
	_materialRequest *MiniAppEntranceMaterialBizOpenDto
}

// NewTaobaoMiniappDistributionMaterialDeleteRequest 初始化TaobaoMiniappDistributionMaterialDeleteAPIRequest对象
func NewTaobaoMiniappDistributionMaterialDeleteRequest() *TaobaoMiniappDistributionMaterialDeleteAPIRequest {
	return &TaobaoMiniappDistributionMaterialDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionMaterialDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.material.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionMaterialDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMaterialRequest is MaterialRequest Setter
// 投放入口素材信息
func (r *TaobaoMiniappDistributionMaterialDeleteAPIRequest) SetMaterialRequest(_materialRequest *MiniAppEntranceMaterialBizOpenDto) error {
	r._materialRequest = _materialRequest
	r.Set("material_request", _materialRequest)
	return nil
}

// GetMaterialRequest MaterialRequest Getter
func (r TaobaoMiniappDistributionMaterialDeleteAPIRequest) GetMaterialRequest() *MiniAppEntranceMaterialBizOpenDto {
	return r._materialRequest
}
