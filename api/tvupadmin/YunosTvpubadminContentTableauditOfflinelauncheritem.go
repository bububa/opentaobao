package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontenttableauditofflinelauncheritem 运营位管理-联盟一体机下线运营位内容
// yunos.tvpubadmin.content.tableaudit.offlinelauncheritem
//
// 运营位管理-联盟一体机下线运营位内容
func Yunostvpubadmincontenttableauditofflinelauncheritem(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontenttableauditofflinelauncheritemAPIRequest, session string) (*tvupadmin.YunostvpubadmincontenttableauditofflinelauncheritemAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontenttableauditofflinelauncheritemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
