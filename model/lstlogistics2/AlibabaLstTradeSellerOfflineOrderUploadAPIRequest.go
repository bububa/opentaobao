package lstlogistics2

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeSellerOfflineOrderUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.offline.order.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeSellerOfflineOrderUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
