package wlb

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoMerchantInventoryAdjustAPIRequest 商家库存调整 API请求
// cainiao.merchant.inventory.adjust
//
// 商家仓库存调整接口，目前仅支持全量更新
type CainiaoMerchantInventoryAdjustAPIRequest struct {
	model.Params
	// 商家仓编辑库存
	_adjustRequest []MerStoreInvAdjustDto
	// 调用方应用名
	_appName string
	// 操作
	_operation string
}

// NewCainiaoMerchantInventoryAdjustRequest 初始化CainiaoMerchantInventoryAdjustAPIRequest对象
func NewCainiaoMerchantInventoryAdjustRequest() *CainiaoMerchantInventoryAdjustAPIRequest {
	return &CainiaoMerchantInventoryAdjustAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoMerchantInventoryAdjustAPIRequest) Reset() {
	r._adjustRequest = r._adjustRequest[:0]
	r._appName = ""
	r._operation = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoMerchantInventoryAdjustAPIRequest) GetApiMethodName() string {
	return "cainiao.merchant.inventory.adjust"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoMerchantInventoryAdjustAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoMerchantInventoryAdjustAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdjustRequest is AdjustRequest Setter
// 商家仓编辑库存
func (r *CainiaoMerchantInventoryAdjustAPIRequest) SetAdjustRequest(_adjustRequest []MerStoreInvAdjustDto) error {
	r._adjustRequest = _adjustRequest
	r.Set("adjust_request", _adjustRequest)
	return nil
}

// GetAdjustRequest AdjustRequest Getter
func (r CainiaoMerchantInventoryAdjustAPIRequest) GetAdjustRequest() []MerStoreInvAdjustDto {
	return r._adjustRequest
}

// SetAppName is AppName Setter
// 调用方应用名
func (r *CainiaoMerchantInventoryAdjustAPIRequest) SetAppName(_appName string) error {
	r._appName = _appName
	r.Set("app_name", _appName)
	return nil
}

// GetAppName AppName Getter
func (r CainiaoMerchantInventoryAdjustAPIRequest) GetAppName() string {
	return r._appName
}

// SetOperation is Operation Setter
// 操作
func (r *CainiaoMerchantInventoryAdjustAPIRequest) SetOperation(_operation string) error {
	r._operation = _operation
	r.Set("operation", _operation)
	return nil
}

// GetOperation Operation Getter
func (r CainiaoMerchantInventoryAdjustAPIRequest) GetOperation() string {
	return r._operation
}

var poolCainiaoMerchantInventoryAdjustAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoMerchantInventoryAdjustRequest()
	},
}

// GetCainiaoMerchantInventoryAdjustRequest 从 sync.Pool 获取 CainiaoMerchantInventoryAdjustAPIRequest
func GetCainiaoMerchantInventoryAdjustAPIRequest() *CainiaoMerchantInventoryAdjustAPIRequest {
	return poolCainiaoMerchantInventoryAdjustAPIRequest.Get().(*CainiaoMerchantInventoryAdjustAPIRequest)
}

// ReleaseCainiaoMerchantInventoryAdjustAPIRequest 将 CainiaoMerchantInventoryAdjustAPIRequest 放入 sync.Pool
func ReleaseCainiaoMerchantInventoryAdjustAPIRequest(v *CainiaoMerchantInventoryAdjustAPIRequest) {
	v.Reset()
	poolCainiaoMerchantInventoryAdjustAPIRequest.Put(v)
}
