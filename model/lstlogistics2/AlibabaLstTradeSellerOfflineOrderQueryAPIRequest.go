package lstlogistics2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttradesellerofflineorderqueryAPIRequest 供应商-线下订单-查询接口 API请求
// alibaba.lst.trade.seller.offline.order.query
//
// 供应商线下订单数据上传后查询物流状态
type AlibabalsttradesellerofflineorderqueryAPIRequest struct {
	model.Params
	// 入参
	_offlineOrderQueryParam *LstOfflineOrderQueryParam
}

// NewAlibabalsttradesellerofflineorderqueryRequest 初始化AlibabalsttradesellerofflineorderqueryAPIRequest对象
func NewAlibabalsttradesellerofflineorderqueryRequest() *AlibabalsttradesellerofflineorderqueryAPIRequest {
	return &AlibabalsttradesellerofflineorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsttradesellerofflineorderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.offline.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsttradesellerofflineorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsttradesellerofflineorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineOrderQueryParam is OfflineOrderQueryParam Setter
// 入参
func (r *AlibabalsttradesellerofflineorderqueryAPIRequest) SetOfflineOrderQueryParam(_offlineOrderQueryParam *LstOfflineOrderQueryParam) error {
	r._offlineOrderQueryParam = _offlineOrderQueryParam
	r.Set("offline_order_query_param", _offlineOrderQueryParam)
	return nil
}

// GetOfflineOrderQueryParam OfflineOrderQueryParam Getter
func (r AlibabalsttradesellerofflineorderqueryAPIRequest) GetOfflineOrderQueryParam() *LstOfflineOrderQueryParam {
	return r._offlineOrderQueryParam
}
