package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentAppOnoffappbylicenseAPIRequest 应用上下架操作 API请求
// yunos.tvpubadmin.content.app.onoffappbylicense
//
// 应用上下架操作
type YunosTvpubadminContentAppOnoffappbylicenseAPIRequest struct {
	model.Params
	// com.ali.yunos.tvacs.domain.OnOffApp
	_onOffApp string
}

// NewYunosTvpubadminContentAppOnoffappbylicenseRequest 初始化YunosTvpubadminContentAppOnoffappbylicenseAPIRequest对象
func NewYunosTvpubadminContentAppOnoffappbylicenseRequest() *YunosTvpubadminContentAppOnoffappbylicenseAPIRequest {
	return &YunosTvpubadminContentAppOnoffappbylicenseAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentAppOnoffappbylicenseAPIRequest) Reset() {
	r._onOffApp = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentAppOnoffappbylicenseAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.app.onoffappbylicense"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentAppOnoffappbylicenseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentAppOnoffappbylicenseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOnOffApp is OnOffApp Setter
// com.ali.yunos.tvacs.domain.OnOffApp
func (r *YunosTvpubadminContentAppOnoffappbylicenseAPIRequest) SetOnOffApp(_onOffApp string) error {
	r._onOffApp = _onOffApp
	r.Set("on_off_app", _onOffApp)
	return nil
}

// GetOnOffApp OnOffApp Getter
func (r YunosTvpubadminContentAppOnoffappbylicenseAPIRequest) GetOnOffApp() string {
	return r._onOffApp
}

var poolYunosTvpubadminContentAppOnoffappbylicenseAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentAppOnoffappbylicenseRequest()
	},
}

// GetYunosTvpubadminContentAppOnoffappbylicenseRequest 从 sync.Pool 获取 YunosTvpubadminContentAppOnoffappbylicenseAPIRequest
func GetYunosTvpubadminContentAppOnoffappbylicenseAPIRequest() *YunosTvpubadminContentAppOnoffappbylicenseAPIRequest {
	return poolYunosTvpubadminContentAppOnoffappbylicenseAPIRequest.Get().(*YunosTvpubadminContentAppOnoffappbylicenseAPIRequest)
}

// ReleaseYunosTvpubadminContentAppOnoffappbylicenseAPIRequest 将 YunosTvpubadminContentAppOnoffappbylicenseAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentAppOnoffappbylicenseAPIRequest(v *YunosTvpubadminContentAppOnoffappbylicenseAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentAppOnoffappbylicenseAPIRequest.Put(v)
}
