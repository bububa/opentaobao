package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsGoodsInfoSyncAPIRequest WMS回传货品长宽高图片等信息 API请求
// taobao.logistics.wms.goods.info.sync
//
// WMS回传货品长宽高图片等信息
type TaobaoLogisticsWmsGoodsInfoSyncAPIRequest struct {
	model.Params
	// 请求
	_wmsGoodsInfoSyncRequest *WmsGoodsInfoSyncRequest
}

// NewTaobaoLogisticsWmsGoodsInfoSyncRequest 初始化TaobaoLogisticsWmsGoodsInfoSyncAPIRequest对象
func NewTaobaoLogisticsWmsGoodsInfoSyncRequest() *TaobaoLogisticsWmsGoodsInfoSyncAPIRequest {
	return &TaobaoLogisticsWmsGoodsInfoSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsWmsGoodsInfoSyncAPIRequest) Reset() {
	r._wmsGoodsInfoSyncRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsWmsGoodsInfoSyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.wms.goods.info.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsWmsGoodsInfoSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsWmsGoodsInfoSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWmsGoodsInfoSyncRequest is WmsGoodsInfoSyncRequest Setter
// 请求
func (r *TaobaoLogisticsWmsGoodsInfoSyncAPIRequest) SetWmsGoodsInfoSyncRequest(_wmsGoodsInfoSyncRequest *WmsGoodsInfoSyncRequest) error {
	r._wmsGoodsInfoSyncRequest = _wmsGoodsInfoSyncRequest
	r.Set("wms_goods_info_sync_request", _wmsGoodsInfoSyncRequest)
	return nil
}

// GetWmsGoodsInfoSyncRequest WmsGoodsInfoSyncRequest Getter
func (r TaobaoLogisticsWmsGoodsInfoSyncAPIRequest) GetWmsGoodsInfoSyncRequest() *WmsGoodsInfoSyncRequest {
	return r._wmsGoodsInfoSyncRequest
}

var poolTaobaoLogisticsWmsGoodsInfoSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsWmsGoodsInfoSyncRequest()
	},
}

// GetTaobaoLogisticsWmsGoodsInfoSyncRequest 从 sync.Pool 获取 TaobaoLogisticsWmsGoodsInfoSyncAPIRequest
func GetTaobaoLogisticsWmsGoodsInfoSyncAPIRequest() *TaobaoLogisticsWmsGoodsInfoSyncAPIRequest {
	return poolTaobaoLogisticsWmsGoodsInfoSyncAPIRequest.Get().(*TaobaoLogisticsWmsGoodsInfoSyncAPIRequest)
}

// ReleaseTaobaoLogisticsWmsGoodsInfoSyncAPIRequest 将 TaobaoLogisticsWmsGoodsInfoSyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsWmsGoodsInfoSyncAPIRequest(v *TaobaoLogisticsWmsGoodsInfoSyncAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsWmsGoodsInfoSyncAPIRequest.Put(v)
}
