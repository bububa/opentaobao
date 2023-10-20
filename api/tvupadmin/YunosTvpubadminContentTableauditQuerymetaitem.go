package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontenttableauditquerymetaitem 运营位管控-查询魔盒运营位元数据列表
// yunos.tvpubadmin.content.tableaudit.querymetaitem
//
// 运营位管控-查询魔盒运营位元数据列表
func Yunostvpubadmincontenttableauditquerymetaitem(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontenttableauditquerymetaitemAPIRequest, session string) (*tvupadmin.YunostvpubadmincontenttableauditquerymetaitemAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontenttableauditquerymetaitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
