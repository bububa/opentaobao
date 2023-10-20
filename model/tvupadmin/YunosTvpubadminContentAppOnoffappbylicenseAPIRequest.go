package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentapponoffappbylicenseAPIRequest 应用上下架操作 API请求
// yunos.tvpubadmin.content.app.onoffappbylicense
//
// 应用上下架操作
type YunostvpubadmincontentapponoffappbylicenseAPIRequest struct {
	model.Params
	// com.ali.yunos.tvacs.domain.OnOffApp
	_onOffApp string
}

// NewYunostvpubadmincontentapponoffappbylicenseRequest 初始化YunostvpubadmincontentapponoffappbylicenseAPIRequest对象
func NewYunostvpubadmincontentapponoffappbylicenseRequest() *YunostvpubadmincontentapponoffappbylicenseAPIRequest {
	return &YunostvpubadmincontentapponoffappbylicenseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentapponoffappbylicenseAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.app.onoffappbylicense"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentapponoffappbylicenseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentapponoffappbylicenseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOnOffApp is OnOffApp Setter
// com.ali.yunos.tvacs.domain.OnOffApp
func (r *YunostvpubadmincontentapponoffappbylicenseAPIRequest) SetOnOffApp(_onOffApp string) error {
	r._onOffApp = _onOffApp
	r.Set("on_off_app", _onOffApp)
	return nil
}

// GetOnOffApp OnOffApp Getter
func (r YunostvpubadmincontentapponoffappbylicenseAPIRequest) GetOnOffApp() string {
	return r._onOffApp
}
