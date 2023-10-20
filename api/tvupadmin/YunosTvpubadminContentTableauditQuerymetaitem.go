package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentTableauditQuerymetaitem 运营位管控-查询魔盒运营位元数据列表
// yunos.tvpubadmin.content.tableaudit.querymetaitem
//
// 运营位管控-查询魔盒运营位元数据列表
func YunosTvpubadminContentTableauditQuerymetaitem(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentTableauditQuerymetaitemAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentTableauditQuerymetaitemAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentTableauditQuerymetaitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
