package wlb

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoBmsOrderConsignConfirmAPIRequest) Reset() {
	r._content = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoBmsOrderConsignConfirmAPIRequest) GetApiMethodName() string {
	return "cainiao.bms.order.consign.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoBmsOrderConsignConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoBmsOrderConsignConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 通知消息主体
func (r *CainiaoBmsOrderConsignConfirmAPIRequest) SetContent(_content *BmsConsignOrderConfirm) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r CainiaoBmsOrderConsignConfirmAPIRequest) GetContent() *BmsConsignOrderConfirm {
	return r._content
}

var poolCainiaoBmsOrderConsignConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoBmsOrderConsignConfirmRequest()
	},
}

// GetCainiaoBmsOrderConsignConfirmRequest 从 sync.Pool 获取 CainiaoBmsOrderConsignConfirmAPIRequest
func GetCainiaoBmsOrderConsignConfirmAPIRequest() *CainiaoBmsOrderConsignConfirmAPIRequest {
	return poolCainiaoBmsOrderConsignConfirmAPIRequest.Get().(*CainiaoBmsOrderConsignConfirmAPIRequest)
}

// ReleaseCainiaoBmsOrderConsignConfirmAPIRequest 将 CainiaoBmsOrderConsignConfirmAPIRequest 放入 sync.Pool
func ReleaseCainiaoBmsOrderConsignConfirmAPIRequest(v *CainiaoBmsOrderConsignConfirmAPIRequest) {
	v.Reset()
	poolCainiaoBmsOrderConsignConfirmAPIRequest.Put(v)
}
