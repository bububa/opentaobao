package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

/* YunosTvpubadminContentAppOnoffappbylicense
应用上下架操作
yunos.tvpubadmin.content.app.onoffappbylicense

应用上下架操作 */
func YunosTvpubadminContentAppOnoffappbylicense(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentAppOnoffappbylicenseAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentAppOnoffappbylicenseAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentAppOnoffappbylicenseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
