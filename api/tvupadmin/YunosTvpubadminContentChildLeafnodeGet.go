package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentchildleafnodeget 获取少儿大厅二级类目列表
// yunos.tvpubadmin.content.child.leafnode.get
//
// 获取少儿大厅二级类目列表的接口
func Yunostvpubadmincontentchildleafnodeget(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentchildleafnodegetAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentchildleafnodegetAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentchildleafnodegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
