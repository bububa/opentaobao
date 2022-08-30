package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionMaterialCreateAPIRequest 小程序投放--新建投放素材 API请求
// taobao.miniapp.distribution.material.create
//
// 为可投放的小程序，增加入口的素材信息，比如图片、引导文案等等。
type TaobaoMiniappDistributionMaterialCreateAPIRequest struct {
	model.Params
	// 投放入口素材信息
	_materialRequest *MiniAppEntranceMaterialBizOpenDto
}

// NewTaobaoMiniappDistributionMaterialCreateRequest 初始化TaobaoMiniappDistributionMaterialCreateAPIRequest对象
func NewTaobaoMiniappDistributionMaterialCreateRequest() *TaobaoMiniappDistributionMaterialCreateAPIRequest {
	return &TaobaoMiniappDistributionMaterialCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionMaterialCreateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.material.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionMaterialCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMaterialRequest is MaterialRequest Setter
// 投放入口素材信息
func (r *TaobaoMiniappDistributionMaterialCreateAPIRequest) SetMaterialRequest(_materialRequest *MiniAppEntranceMaterialBizOpenDto) error {
	r._materialRequest = _materialRequest
	r.Set("material_request", _materialRequest)
	return nil
}

// GetMaterialRequest MaterialRequest Getter
func (r TaobaoMiniappDistributionMaterialCreateAPIRequest) GetMaterialRequest() *MiniAppEntranceMaterialBizOpenDto {
	return r._materialRequest
}
