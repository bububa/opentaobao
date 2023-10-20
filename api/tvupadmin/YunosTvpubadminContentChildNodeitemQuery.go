package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentchildnodeitemquery 查询少儿大厅类目内容
// yunos.tvpubadmin.content.child.nodeitem.query
//
// 查询少儿大厅类目内容信息
func Yunostvpubadmincontentchildnodeitemquery(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentchildnodeitemqueryAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentchildnodeitemqueryAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentchildnodeitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
