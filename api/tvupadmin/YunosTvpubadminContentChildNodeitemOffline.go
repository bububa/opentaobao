package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentchildnodeitemoffline 少儿大厅类目内容下线接口
// yunos.tvpubadmin.content.child.nodeitem.offline
//
// 少儿大厅类目内容下线接口
func Yunostvpubadmincontentchildnodeitemoffline(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentchildnodeitemofflineAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentchildnodeitemofflineAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentchildnodeitemofflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
