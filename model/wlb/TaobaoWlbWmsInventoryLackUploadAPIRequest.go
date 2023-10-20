package wlb

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsInventoryLackUploadAPIRequest 缺货通知 API请求
// taobao.wlb.wms.inventory.lack.upload
//
// 缺货通知
type TaobaoWlbWmsInventoryLackUploadAPIRequest struct {
	model.Params
	// 缺货通知信息
	_content *WlbWmsInventoryLackUpload
}

// NewTaobaoWlbWmsInventoryLackUploadRequest 初始化TaobaoWlbWmsInventoryLackUploadAPIRequest对象
func NewTaobaoWlbWmsInventoryLackUploadRequest() *TaobaoWlbWmsInventoryLackUploadAPIRequest {
	return &TaobaoWlbWmsInventoryLackUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbWmsInventoryLackUploadAPIRequest) Reset() {
	r._content = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsInventoryLackUploadAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.inventory.lack.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsInventoryLackUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbWmsInventoryLackUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 缺货通知信息
func (r *TaobaoWlbWmsInventoryLackUploadAPIRequest) SetContent(_content *WlbWmsInventoryLackUpload) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoWlbWmsInventoryLackUploadAPIRequest) GetContent() *WlbWmsInventoryLackUpload {
	return r._content
}

var poolTaobaoWlbWmsInventoryLackUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbWmsInventoryLackUploadRequest()
	},
}

// GetTaobaoWlbWmsInventoryLackUploadRequest 从 sync.Pool 获取 TaobaoWlbWmsInventoryLackUploadAPIRequest
func GetTaobaoWlbWmsInventoryLackUploadAPIRequest() *TaobaoWlbWmsInventoryLackUploadAPIRequest {
	return poolTaobaoWlbWmsInventoryLackUploadAPIRequest.Get().(*TaobaoWlbWmsInventoryLackUploadAPIRequest)
}

// ReleaseTaobaoWlbWmsInventoryLackUploadAPIRequest 将 TaobaoWlbWmsInventoryLackUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbWmsInventoryLackUploadAPIRequest(v *TaobaoWlbWmsInventoryLackUploadAPIRequest) {
	v.Reset()
	poolTaobaoWlbWmsInventoryLackUploadAPIRequest.Put(v)
}
