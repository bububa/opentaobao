package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentappquerybylicence 按牌照查询应用
// yunos.tvpubadmin.content.app.querybylicence
//
// 按牌照查询应用
func Yunostvpubadmincontentappquerybylicence(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentappquerybylicenceAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentappquerybylicenceAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentappquerybylicenceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
