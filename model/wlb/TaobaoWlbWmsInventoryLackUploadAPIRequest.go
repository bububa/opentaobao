package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwmsinventorylackuploadAPIRequest 缺货通知 API请求
// taobao.wlb.wms.inventory.lack.upload
//
// 缺货通知
type TaobaowlbwmsinventorylackuploadAPIRequest struct {
	model.Params
	// 缺货通知信息
	_content *WlbWmsInventoryLackUpload
}

// NewTaobaowlbwmsinventorylackuploadRequest 初始化TaobaowlbwmsinventorylackuploadAPIRequest对象
func NewTaobaowlbwmsinventorylackuploadRequest() *TaobaowlbwmsinventorylackuploadAPIRequest {
	return &TaobaowlbwmsinventorylackuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwmsinventorylackuploadAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.inventory.lack.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwmsinventorylackuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwmsinventorylackuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 缺货通知信息
func (r *TaobaowlbwmsinventorylackuploadAPIRequest) SetContent(_content *WlbWmsInventoryLackUpload) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaowlbwmsinventorylackuploadAPIRequest) GetContent() *WlbWmsInventoryLackUpload {
	return r._content
}
