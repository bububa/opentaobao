package miniappopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionMaterialUpdateAPIRequest 小程序投放 --- 更新投放素材 API请求
// taobao.miniapp.distribution.material.update
//
// 修改已录入的投放素材信息。
type TaobaoMiniappDistributionMaterialUpdateAPIRequest struct {
	model.Params
	// 投放入口素材信息
	_materialRequest *MiniAppEntranceMaterialBizOpenDto
}

// NewTaobaoMiniappDistributionMaterialUpdateRequest 初始化TaobaoMiniappDistributionMaterialUpdateAPIRequest对象
func NewTaobaoMiniappDistributionMaterialUpdateRequest() *TaobaoMiniappDistributionMaterialUpdateAPIRequest {
	return &TaobaoMiniappDistributionMaterialUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappDistributionMaterialUpdateAPIRequest) Reset() {
	r._materialRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionMaterialUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.material.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionMaterialUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappDistributionMaterialUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMaterialRequest is MaterialRequest Setter
// 投放入口素材信息
func (r *TaobaoMiniappDistributionMaterialUpdateAPIRequest) SetMaterialRequest(_materialRequest *MiniAppEntranceMaterialBizOpenDto) error {
	r._materialRequest = _materialRequest
	r.Set("material_request", _materialRequest)
	return nil
}

// GetMaterialRequest MaterialRequest Getter
func (r TaobaoMiniappDistributionMaterialUpdateAPIRequest) GetMaterialRequest() *MiniAppEntranceMaterialBizOpenDto {
	return r._materialRequest
}

var poolTaobaoMiniappDistributionMaterialUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappDistributionMaterialUpdateRequest()
	},
}

// GetTaobaoMiniappDistributionMaterialUpdateRequest 从 sync.Pool 获取 TaobaoMiniappDistributionMaterialUpdateAPIRequest
func GetTaobaoMiniappDistributionMaterialUpdateAPIRequest() *TaobaoMiniappDistributionMaterialUpdateAPIRequest {
	return poolTaobaoMiniappDistributionMaterialUpdateAPIRequest.Get().(*TaobaoMiniappDistributionMaterialUpdateAPIRequest)
}

// ReleaseTaobaoMiniappDistributionMaterialUpdateAPIRequest 将 TaobaoMiniappDistributionMaterialUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappDistributionMaterialUpdateAPIRequest(v *TaobaoMiniappDistributionMaterialUpdateAPIRequest) {
	v.Reset()
	poolTaobaoMiniappDistributionMaterialUpdateAPIRequest.Put(v)
}
