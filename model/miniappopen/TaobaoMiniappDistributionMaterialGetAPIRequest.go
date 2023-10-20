package miniappopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionMaterialGetAPIRequest 小程序投放---读取投放入口素材信息 API请求
// taobao.miniapp.distribution.material.get
//
// 读取已录入的投放入口素材信息。
type TaobaoMiniappDistributionMaterialGetAPIRequest struct {
	model.Params
	// 投放入口素材信息
	_materialRequest *MiniAppEntranceMaterialBizOpenDto
}

// NewTaobaoMiniappDistributionMaterialGetRequest 初始化TaobaoMiniappDistributionMaterialGetAPIRequest对象
func NewTaobaoMiniappDistributionMaterialGetRequest() *TaobaoMiniappDistributionMaterialGetAPIRequest {
	return &TaobaoMiniappDistributionMaterialGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappDistributionMaterialGetAPIRequest) Reset() {
	r._materialRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionMaterialGetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.material.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionMaterialGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappDistributionMaterialGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMaterialRequest is MaterialRequest Setter
// 投放入口素材信息
func (r *TaobaoMiniappDistributionMaterialGetAPIRequest) SetMaterialRequest(_materialRequest *MiniAppEntranceMaterialBizOpenDto) error {
	r._materialRequest = _materialRequest
	r.Set("material_request", _materialRequest)
	return nil
}

// GetMaterialRequest MaterialRequest Getter
func (r TaobaoMiniappDistributionMaterialGetAPIRequest) GetMaterialRequest() *MiniAppEntranceMaterialBizOpenDto {
	return r._materialRequest
}

var poolTaobaoMiniappDistributionMaterialGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappDistributionMaterialGetRequest()
	},
}

// GetTaobaoMiniappDistributionMaterialGetRequest 从 sync.Pool 获取 TaobaoMiniappDistributionMaterialGetAPIRequest
func GetTaobaoMiniappDistributionMaterialGetAPIRequest() *TaobaoMiniappDistributionMaterialGetAPIRequest {
	return poolTaobaoMiniappDistributionMaterialGetAPIRequest.Get().(*TaobaoMiniappDistributionMaterialGetAPIRequest)
}

// ReleaseTaobaoMiniappDistributionMaterialGetAPIRequest 将 TaobaoMiniappDistributionMaterialGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappDistributionMaterialGetAPIRequest(v *TaobaoMiniappDistributionMaterialGetAPIRequest) {
	v.Reset()
	poolTaobaoMiniappDistributionMaterialGetAPIRequest.Put(v)
}
