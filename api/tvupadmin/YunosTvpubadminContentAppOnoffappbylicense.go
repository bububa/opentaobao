package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentapponoffappbylicense 应用上下架操作
// yunos.tvpubadmin.content.app.onoffappbylicense
//
// 应用上下架操作
func Yunostvpubadmincontentapponoffappbylicense(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentapponoffappbylicenseAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentapponoffappbylicenseAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentapponoffappbylicenseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
