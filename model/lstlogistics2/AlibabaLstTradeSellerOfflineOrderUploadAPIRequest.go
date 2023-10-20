package lstlogistics2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeSellerOfflineOrderUploadAPIRequest 供应商-线下订单-导入接口 API请求
// alibaba.lst.trade.seller.offline.order.upload
//
// 供应商线下订单数据上传、实现和零售通本地云仓订单的共配
type AlibabaLstTradeSellerOfflineOrderUploadAPIRequest struct {
	model.Params
	// 入参
	_offlineOrderUploadParam *LstOffLineOrderUploadParam
}

// NewAlibabaLstTradeSellerOfflineOrderUploadRequest 初始化AlibabaLstTradeSellerOfflineOrderUploadAPIRequest对象
func NewAlibabaLstTradeSellerOfflineOrderUploadRequest() *AlibabaLstTradeSellerOfflineOrderUploadAPIRequest {
	return &AlibabaLstTradeSellerOfflineOrderUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstTradeSellerOfflineOrderUploadAPIRequest) Reset() {
	r._offlineOrderUploadParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeSellerOfflineOrderUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.offline.order.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeSellerOfflineOrderUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstTradeSellerOfflineOrderUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineOrderUploadParam is OfflineOrderUploadParam Setter
// 入参
func (r *AlibabaLstTradeSellerOfflineOrderUploadAPIRequest) SetOfflineOrderUploadParam(_offlineOrderUploadParam *LstOffLineOrderUploadParam) error {
	r._offlineOrderUploadParam = _offlineOrderUploadParam
	r.Set("offline_order_upload_param", _offlineOrderUploadParam)
	return nil
}

// GetOfflineOrderUploadParam OfflineOrderUploadParam Getter
func (r AlibabaLstTradeSellerOfflineOrderUploadAPIRequest) GetOfflineOrderUploadParam() *LstOffLineOrderUploadParam {
	return r._offlineOrderUploadParam
}

var poolAlibabaLstTradeSellerOfflineOrderUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstTradeSellerOfflineOrderUploadRequest()
	},
}

// GetAlibabaLstTradeSellerOfflineOrderUploadRequest 从 sync.Pool 获取 AlibabaLstTradeSellerOfflineOrderUploadAPIRequest
func GetAlibabaLstTradeSellerOfflineOrderUploadAPIRequest() *AlibabaLstTradeSellerOfflineOrderUploadAPIRequest {
	return poolAlibabaLstTradeSellerOfflineOrderUploadAPIRequest.Get().(*AlibabaLstTradeSellerOfflineOrderUploadAPIRequest)
}

// ReleaseAlibabaLstTradeSellerOfflineOrderUploadAPIRequest 将 AlibabaLstTradeSellerOfflineOrderUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstTradeSellerOfflineOrderUploadAPIRequest(v *AlibabaLstTradeSellerOfflineOrderUploadAPIRequest) {
	v.Reset()
	poolAlibabaLstTradeSellerOfflineOrderUploadAPIRequest.Put(v)
}
