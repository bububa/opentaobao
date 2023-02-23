package lstlogistics2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeSellerOfflineOrderQueryAPIRequest 供应商-线下订单-查询接口 API请求
// alibaba.lst.trade.seller.offline.order.query
//
// 供应商线下订单数据上传后查询物流状态
type AlibabaLstTradeSellerOfflineOrderQueryAPIRequest struct {
	model.Params
	// 入参
	_offlineOrderQueryParam *LstOfflineOrderQueryParam
}

// NewAlibabaLstTradeSellerOfflineOrderQueryRequest 初始化AlibabaLstTradeSellerOfflineOrderQueryAPIRequest对象
func NewAlibabaLstTradeSellerOfflineOrderQueryRequest() *AlibabaLstTradeSellerOfflineOrderQueryAPIRequest {
	return &AlibabaLstTradeSellerOfflineOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeSellerOfflineOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.offline.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeSellerOfflineOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstTradeSellerOfflineOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineOrderQueryParam is OfflineOrderQueryParam Setter
// 入参
func (r *AlibabaLstTradeSellerOfflineOrderQueryAPIRequest) SetOfflineOrderQueryParam(_offlineOrderQueryParam *LstOfflineOrderQueryParam) error {
	r._offlineOrderQueryParam = _offlineOrderQueryParam
	r.Set("offline_order_query_param", _offlineOrderQueryParam)
	return nil
}

// GetOfflineOrderQueryParam OfflineOrderQueryParam Getter
func (r AlibabaLstTradeSellerOfflineOrderQueryAPIRequest) GetOfflineOrderQueryParam() *LstOfflineOrderQueryParam {
	return r._offlineOrderQueryParam
}
