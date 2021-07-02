package tvupadmin

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentAppOnoffappbylicenseAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.app.onoffappbylicense"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentAppOnoffappbylicenseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
