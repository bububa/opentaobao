package miniappopen

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappDistributionMaterialDeleteAPIRequest) Reset() {
	r._materialRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionMaterialDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.material.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionMaterialDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappDistributionMaterialDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoMiniappDistributionMaterialDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappDistributionMaterialDeleteRequest()
	},
}

// GetTaobaoMiniappDistributionMaterialDeleteRequest 从 sync.Pool 获取 TaobaoMiniappDistributionMaterialDeleteAPIRequest
func GetTaobaoMiniappDistributionMaterialDeleteAPIRequest() *TaobaoMiniappDistributionMaterialDeleteAPIRequest {
	return poolTaobaoMiniappDistributionMaterialDeleteAPIRequest.Get().(*TaobaoMiniappDistributionMaterialDeleteAPIRequest)
}

// ReleaseTaobaoMiniappDistributionMaterialDeleteAPIRequest 将 TaobaoMiniappDistributionMaterialDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappDistributionMaterialDeleteAPIRequest(v *TaobaoMiniappDistributionMaterialDeleteAPIRequest) {
	v.Reset()
	poolTaobaoMiniappDistributionMaterialDeleteAPIRequest.Put(v)
}
