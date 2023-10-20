package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontenttableauditquerylauncher 运营位管控-查询联盟一体机运营位元数据列表
// yunos.tvpubadmin.content.tableaudit.querylauncher
//
// 运营位管控-查询联盟一体机运营位元数据列表
func Yunostvpubadmincontenttableauditquerylauncher(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontenttableauditquerylauncherAPIRequest, session string) (*tvupadmin.YunostvpubadmincontenttableauditquerylauncherAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontenttableauditquerylauncherAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
