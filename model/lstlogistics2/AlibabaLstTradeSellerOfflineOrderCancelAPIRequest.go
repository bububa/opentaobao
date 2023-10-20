package lstlogistics2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeSellerOfflineOrderCancelAPIRequest 供应商-线下订单-取消接口 API请求
// alibaba.lst.trade.seller.offline.order.cancel
//
// 供应商线下订单数据上传之后取消
type AlibabaLstTradeSellerOfflineOrderCancelAPIRequest struct {
	model.Params
	// 入参
	_offlineOrderCancalParam *LstOfflineOrderCancalParam
}

// NewAlibabaLstTradeSellerOfflineOrderCancelRequest 初始化AlibabaLstTradeSellerOfflineOrderCancelAPIRequest对象
func NewAlibabaLstTradeSellerOfflineOrderCancelRequest() *AlibabaLstTradeSellerOfflineOrderCancelAPIRequest {
	return &AlibabaLstTradeSellerOfflineOrderCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstTradeSellerOfflineOrderCancelAPIRequest) Reset() {
	r._offlineOrderCancalParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeSellerOfflineOrderCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.offline.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeSellerOfflineOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstTradeSellerOfflineOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineOrderCancalParam is OfflineOrderCancalParam Setter
// 入参
func (r *AlibabaLstTradeSellerOfflineOrderCancelAPIRequest) SetOfflineOrderCancalParam(_offlineOrderCancalParam *LstOfflineOrderCancalParam) error {
	r._offlineOrderCancalParam = _offlineOrderCancalParam
	r.Set("offline_order_cancal_param", _offlineOrderCancalParam)
	return nil
}

// GetOfflineOrderCancalParam OfflineOrderCancalParam Getter
func (r AlibabaLstTradeSellerOfflineOrderCancelAPIRequest) GetOfflineOrderCancalParam() *LstOfflineOrderCancalParam {
	return r._offlineOrderCancalParam
}

var poolAlibabaLstTradeSellerOfflineOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstTradeSellerOfflineOrderCancelRequest()
	},
}

// GetAlibabaLstTradeSellerOfflineOrderCancelRequest 从 sync.Pool 获取 AlibabaLstTradeSellerOfflineOrderCancelAPIRequest
func GetAlibabaLstTradeSellerOfflineOrderCancelAPIRequest() *AlibabaLstTradeSellerOfflineOrderCancelAPIRequest {
	return poolAlibabaLstTradeSellerOfflineOrderCancelAPIRequest.Get().(*AlibabaLstTradeSellerOfflineOrderCancelAPIRequest)
}

// ReleaseAlibabaLstTradeSellerOfflineOrderCancelAPIRequest 将 AlibabaLstTradeSellerOfflineOrderCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstTradeSellerOfflineOrderCancelAPIRequest(v *AlibabaLstTradeSellerOfflineOrderCancelAPIRequest) {
	v.Reset()
	poolAlibabaLstTradeSellerOfflineOrderCancelAPIRequest.Put(v)
}
