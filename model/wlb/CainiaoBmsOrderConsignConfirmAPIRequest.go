package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaobmsorderconsignconfirmAPIRequest BMS出库通知 API请求
// cainiao.bms.order.consign.confirm
//
// BMS出库后，通知ISV
type CainiaobmsorderconsignconfirmAPIRequest struct {
	model.Params
	// 通知消息主体
	_content *BmsConsignOrderConfirm
}

// NewCainiaobmsorderconsignconfirmRequest 初始化CainiaobmsorderconsignconfirmAPIRequest对象
func NewCainiaobmsorderconsignconfirmRequest() *CainiaobmsorderconsignconfirmAPIRequest {
	return &CainiaobmsorderconsignconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaobmsorderconsignconfirmAPIRequest) GetApiMethodName() string {
	return "cainiao.bms.order.consign.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaobmsorderconsignconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaobmsorderconsignconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 通知消息主体
func (r *CainiaobmsorderconsignconfirmAPIRequest) SetContent(_content *BmsConsignOrderConfirm) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r CainiaobmsorderconsignconfirmAPIRequest) GetContent() *BmsConsignOrderConfirm {
	return r._content
}
