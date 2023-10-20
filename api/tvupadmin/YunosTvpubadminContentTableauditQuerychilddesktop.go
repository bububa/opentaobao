package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontenttableauditquerychilddesktop 迎客松查看小酷宝桌面坑位元数据列表
// yunos.tvpubadmin.content.tableaudit.querychilddesktop
//
// 迎客松查看小酷宝桌面坑位元数据列表
func Yunostvpubadmincontenttableauditquerychilddesktop(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontenttableauditquerychilddesktopAPIRequest, session string) (*tvupadmin.YunostvpubadmincontenttableauditquerychilddesktopAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontenttableauditquerychilddesktopAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
