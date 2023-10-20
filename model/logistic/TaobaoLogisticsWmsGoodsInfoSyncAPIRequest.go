package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswmsgoodsinfosyncAPIRequest WMS回传货品长宽高图片等信息 API请求
// taobao.logistics.wms.goods.info.sync
//
// WMS回传货品长宽高图片等信息
type TaobaologisticswmsgoodsinfosyncAPIRequest struct {
	model.Params
	// 请求
	_wmsGoodsInfoSyncRequest *WmsGoodsInfoSyncRequest
}

// NewTaobaologisticswmsgoodsinfosyncRequest 初始化TaobaologisticswmsgoodsinfosyncAPIRequest对象
func NewTaobaologisticswmsgoodsinfosyncRequest() *TaobaologisticswmsgoodsinfosyncAPIRequest {
	return &TaobaologisticswmsgoodsinfosyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticswmsgoodsinfosyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.wms.goods.info.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticswmsgoodsinfosyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticswmsgoodsinfosyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWmsGoodsInfoSyncRequest is WmsGoodsInfoSyncRequest Setter
// 请求
func (r *TaobaologisticswmsgoodsinfosyncAPIRequest) SetWmsGoodsInfoSyncRequest(_wmsGoodsInfoSyncRequest *WmsGoodsInfoSyncRequest) error {
	r._wmsGoodsInfoSyncRequest = _wmsGoodsInfoSyncRequest
	r.Set("wms_goods_info_sync_request", _wmsGoodsInfoSyncRequest)
	return nil
}

// GetWmsGoodsInfoSyncRequest WmsGoodsInfoSyncRequest Getter
func (r TaobaologisticswmsgoodsinfosyncAPIRequest) GetWmsGoodsInfoSyncRequest() *WmsGoodsInfoSyncRequest {
	return r._wmsGoodsInfoSyncRequest
}
