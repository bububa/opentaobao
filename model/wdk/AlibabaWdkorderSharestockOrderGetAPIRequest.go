package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkordersharestockordergetAPIRequest 猫超商户订单拉取 API请求
// alibaba.wdkorder.sharestock.order.get
//
// 商户拉取猫超订单数据
type AlibabawdkordersharestockordergetAPIRequest struct {
	model.Params
	// 淘宝主订单ID
	_tbOrderId int64
}

// NewAlibabawdkordersharestockordergetRequest 初始化AlibabawdkordersharestockordergetAPIRequest对象
func NewAlibabawdkordersharestockordergetRequest() *AlibabawdkordersharestockordergetAPIRequest {
	return &AlibabawdkordersharestockordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkordersharestockordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkordersharestockordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkordersharestockordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTbOrderId is TbOrderId Setter
// 淘宝主订单ID
func (r *AlibabawdkordersharestockordergetAPIRequest) SetTbOrderId(_tbOrderId int64) error {
	r._tbOrderId = _tbOrderId
	r.Set("tb_order_id", _tbOrderId)
	return nil
}

// GetTbOrderId TbOrderId Getter
func (r AlibabawdkordersharestockordergetAPIRequest) GetTbOrderId() int64 {
	return r._tbOrderId
}
