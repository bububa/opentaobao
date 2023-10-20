package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkordersharestockinsurancegetorderAPIRequest 共享库存订单投保消息获取 API请求
// alibaba.wdkorder.sharestock.insurance.getorder
//
// 共享库存订单投保消息获取
type AlibabawdkordersharestockinsurancegetorderAPIRequest struct {
	model.Params
	// 淘宝子订单ID
	_tbSubOrderId int64
}

// NewAlibabawdkordersharestockinsurancegetorderRequest 初始化AlibabawdkordersharestockinsurancegetorderAPIRequest对象
func NewAlibabawdkordersharestockinsurancegetorderRequest() *AlibabawdkordersharestockinsurancegetorderAPIRequest {
	return &AlibabawdkordersharestockinsurancegetorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkordersharestockinsurancegetorderAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.insurance.getorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkordersharestockinsurancegetorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkordersharestockinsurancegetorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTbSubOrderId is TbSubOrderId Setter
// 淘宝子订单ID
func (r *AlibabawdkordersharestockinsurancegetorderAPIRequest) SetTbSubOrderId(_tbSubOrderId int64) error {
	r._tbSubOrderId = _tbSubOrderId
	r.Set("tb_sub_order_id", _tbSubOrderId)
	return nil
}

// GetTbSubOrderId TbSubOrderId Getter
func (r AlibabawdkordersharestockinsurancegetorderAPIRequest) GetTbSubOrderId() int64 {
	return r._tbSubOrderId
}
