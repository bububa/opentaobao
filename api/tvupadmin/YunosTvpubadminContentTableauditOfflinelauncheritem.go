package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentTableauditOfflinelauncheritem 运营位管理-联盟一体机下线运营位内容
// yunos.tvpubadmin.content.tableaudit.offlinelauncheritem
//
// 运营位管理-联盟一体机下线运营位内容
func YunosTvpubadminContentTableauditOfflinelauncheritem(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
