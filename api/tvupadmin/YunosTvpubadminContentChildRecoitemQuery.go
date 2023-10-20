package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentchildrecoitemquery 查询少儿大厅推荐内容列表
// yunos.tvpubadmin.content.child.recoitem.query
//
// 查询少儿大厅推荐内容列表
func Yunostvpubadmincontentchildrecoitemquery(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentchildrecoitemqueryAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentchildrecoitemqueryAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentchildrecoitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
