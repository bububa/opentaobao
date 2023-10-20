package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBmPaiyangSupplyPublishAPIRequest 派样商品库存变更同步接口 API请求
// alibaba.wdk.bm.paiyang.supply.publish
//
// 淘鲜达接入第三方进行派样，第三方同步大仓和门店的库存变更信息。
type AlibabaWdkBmPaiyangSupplyPublishAPIRequest struct {
	model.Params
	// 请求入参
	_isvSupplySyncParam *IsvSupplySyncParam
}

// NewAlibabaWdkBmPaiyangSupplyPublishRequest 初始化AlibabaWdkBmPaiyangSupplyPublishAPIRequest对象
func NewAlibabaWdkBmPaiyangSupplyPublishRequest() *AlibabaWdkBmPaiyangSupplyPublishAPIRequest {
	return &AlibabaWdkBmPaiyangSupplyPublishAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkBmPaiyangSupplyPublishAPIRequest) Reset() {
	r._isvSupplySyncParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmPaiyangSupplyPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bm.paiyang.supply.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmPaiyangSupplyPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkBmPaiyangSupplyPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvSupplySyncParam is IsvSupplySyncParam Setter
// 请求入参
func (r *AlibabaWdkBmPaiyangSupplyPublishAPIRequest) SetIsvSupplySyncParam(_isvSupplySyncParam *IsvSupplySyncParam) error {
	r._isvSupplySyncParam = _isvSupplySyncParam
	r.Set("isv_supply_sync_param", _isvSupplySyncParam)
	return nil
}

// GetIsvSupplySyncParam IsvSupplySyncParam Getter
func (r AlibabaWdkBmPaiyangSupplyPublishAPIRequest) GetIsvSupplySyncParam() *IsvSupplySyncParam {
	return r._isvSupplySyncParam
}

var poolAlibabaWdkBmPaiyangSupplyPublishAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkBmPaiyangSupplyPublishRequest()
	},
}

// GetAlibabaWdkBmPaiyangSupplyPublishRequest 从 sync.Pool 获取 AlibabaWdkBmPaiyangSupplyPublishAPIRequest
func GetAlibabaWdkBmPaiyangSupplyPublishAPIRequest() *AlibabaWdkBmPaiyangSupplyPublishAPIRequest {
	return poolAlibabaWdkBmPaiyangSupplyPublishAPIRequest.Get().(*AlibabaWdkBmPaiyangSupplyPublishAPIRequest)
}

// ReleaseAlibabaWdkBmPaiyangSupplyPublishAPIRequest 将 AlibabaWdkBmPaiyangSupplyPublishAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkBmPaiyangSupplyPublishAPIRequest(v *AlibabaWdkBmPaiyangSupplyPublishAPIRequest) {
	v.Reset()
	poolAlibabaWdkBmPaiyangSupplyPublishAPIRequest.Put(v)
}
