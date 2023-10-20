package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantorderbatchuploadAPIRequest 商家订单数据批量上传 API请求
// alibaba.tcls.aelophy.merchant.order.batch.upload
//
// 商家订单数据上传
type AlibabatclsaelophymerchantorderbatchuploadAPIRequest struct {
	model.Params
	// 商家订单信息
	_orderInfoList []MerchantOrderInfo
}

// NewAlibabatclsaelophymerchantorderbatchuploadRequest 初始化AlibabatclsaelophymerchantorderbatchuploadAPIRequest对象
func NewAlibabatclsaelophymerchantorderbatchuploadRequest() *AlibabatclsaelophymerchantorderbatchuploadAPIRequest {
	return &AlibabatclsaelophymerchantorderbatchuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophymerchantorderbatchuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.order.batch.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophymerchantorderbatchuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophymerchantorderbatchuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderInfoList is OrderInfoList Setter
// 商家订单信息
func (r *AlibabatclsaelophymerchantorderbatchuploadAPIRequest) SetOrderInfoList(_orderInfoList []MerchantOrderInfo) error {
	r._orderInfoList = _orderInfoList
	r.Set("order_info_list", _orderInfoList)
	return nil
}

// GetOrderInfoList OrderInfoList Getter
func (r AlibabatclsaelophymerchantorderbatchuploadAPIRequest) GetOrderInfoList() []MerchantOrderInfo {
	return r._orderInfoList
}
