package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOldposOrderCreateAPIRequest 淘鲜达外部商户老pos机产生的订单同步进淘鲜达 API请求
// alibaba.wdk.oldpos.order.create
//
// 淘鲜达外部商户老pos机产生的订单同步进淘鲜达
type AlibabaWdkOldposOrderCreateAPIRequest struct {
	model.Params
	// 入参
	_posOrderCreateRequest *PosOrderCreateRequest
}

// NewAlibabaWdkOldposOrderCreateRequest 初始化AlibabaWdkOldposOrderCreateAPIRequest对象
func NewAlibabaWdkOldposOrderCreateRequest() *AlibabaWdkOldposOrderCreateAPIRequest {
	return &AlibabaWdkOldposOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkOldposOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.oldpos.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkOldposOrderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PosOrderCreateRequest Setter
// 入参
func (r *AlibabaWdkOldposOrderCreateAPIRequest) SetPosOrderCreateRequest(_posOrderCreateRequest *PosOrderCreateRequest) error {
	r._posOrderCreateRequest = _posOrderCreateRequest
	r.Set("pos_order_create_request", _posOrderCreateRequest)
	return nil
}

// Get PosOrderCreateRequest Getter
func (r AlibabaWdkOldposOrderCreateAPIRequest) GetPosOrderCreateRequest() *PosOrderCreateRequest {
	return r._posOrderCreateRequest
}
