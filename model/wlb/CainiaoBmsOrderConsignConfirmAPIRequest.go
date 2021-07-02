package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoBmsOrderConsignConfirmAPIRequest BMS出库通知 API请求
// cainiao.bms.order.consign.confirm
//
// BMS出库后，通知ISV
type CainiaoBmsOrderConsignConfirmAPIRequest struct {
	model.Params
	// 通知消息主体
	_content *BmsConsignOrderConfirm
}

// NewCainiaoBmsOrderConsignConfirmRequest 初始化CainiaoBmsOrderConsignConfirmAPIRequest对象
func NewCainiaoBmsOrderConsignConfirmRequest() *CainiaoBmsOrderConsignConfirmAPIRequest {
	return &CainiaoBmsOrderConsignConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoBmsOrderConsignConfirmAPIRequest) GetApiMethodName() string {
	return "cainiao.bms.order.consign.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoBmsOrderConsignConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Content Setter
// 通知消息主体
func (r *CainiaoBmsOrderConsignConfirmAPIRequest) SetContent(_content *BmsConsignOrderConfirm) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// Get Content Getter
func (r CainiaoBmsOrderConsignConfirmAPIRequest) GetContent() *BmsConsignOrderConfirm {
	return r._content
}
