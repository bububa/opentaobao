package lstlogistics2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttradesellerofflineordercancelAPIRequest 供应商-线下订单-取消接口 API请求
// alibaba.lst.trade.seller.offline.order.cancel
//
// 供应商线下订单数据上传之后取消
type AlibabalsttradesellerofflineordercancelAPIRequest struct {
	model.Params
	// 入参
	_offlineOrderCancalParam *LstOfflineOrderCancalParam
}

// NewAlibabalsttradesellerofflineordercancelRequest 初始化AlibabalsttradesellerofflineordercancelAPIRequest对象
func NewAlibabalsttradesellerofflineordercancelRequest() *AlibabalsttradesellerofflineordercancelAPIRequest {
	return &AlibabalsttradesellerofflineordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsttradesellerofflineordercancelAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.offline.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsttradesellerofflineordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsttradesellerofflineordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineOrderCancalParam is OfflineOrderCancalParam Setter
// 入参
func (r *AlibabalsttradesellerofflineordercancelAPIRequest) SetOfflineOrderCancalParam(_offlineOrderCancalParam *LstOfflineOrderCancalParam) error {
	r._offlineOrderCancalParam = _offlineOrderCancalParam
	r.Set("offline_order_cancal_param", _offlineOrderCancalParam)
	return nil
}

// GetOfflineOrderCancalParam OfflineOrderCancalParam Getter
func (r AlibabalsttradesellerofflineordercancelAPIRequest) GetOfflineOrderCancalParam() *LstOfflineOrderCancalParam {
	return r._offlineOrderCancalParam
}
