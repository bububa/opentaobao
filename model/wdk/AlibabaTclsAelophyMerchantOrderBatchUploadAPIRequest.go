package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantOrderBatchUploadAPIRequest 商家订单数据批量上传 API请求
// alibaba.tcls.aelophy.merchant.order.batch.upload
//
// 商家订单数据上传
type AlibabaTclsAelophyMerchantOrderBatchUploadAPIRequest struct {
	model.Params
	// 商家订单信息
	_orderInfoList []MerchantOrderInfo
}

// NewAlibabaTclsAelophyMerchantOrderBatchUploadRequest 初始化AlibabaTclsAelophyMerchantOrderBatchUploadAPIRequest对象
func NewAlibabaTclsAelophyMerchantOrderBatchUploadRequest() *AlibabaTclsAelophyMerchantOrderBatchUploadAPIRequest {
	return &AlibabaTclsAelophyMerchantOrderBatchUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantOrderBatchUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.order.batch.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantOrderBatchUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyMerchantOrderBatchUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderInfoList is OrderInfoList Setter
// 商家订单信息
func (r *AlibabaTclsAelophyMerchantOrderBatchUploadAPIRequest) SetOrderInfoList(_orderInfoList []MerchantOrderInfo) error {
	r._orderInfoList = _orderInfoList
	r.Set("order_info_list", _orderInfoList)
	return nil
}

// GetOrderInfoList OrderInfoList Getter
func (r AlibabaTclsAelophyMerchantOrderBatchUploadAPIRequest) GetOrderInfoList() []MerchantOrderInfo {
	return r._orderInfoList
}
