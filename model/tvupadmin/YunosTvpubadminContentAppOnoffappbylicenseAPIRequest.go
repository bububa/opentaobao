package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentAppOnoffappbylicenseAPIRequest
应用上下架操作 API请求
yunos.tvpubadmin.content.app.onoffappbylicense

应用上下架操作 */
type YunosTvpubadminContentAppOnoffappbylicenseAPIRequest struct {
	model.Params
	// com.ali.yunos.tvacs.domain.OnOffApp
	_onOffApp string
}

// New
