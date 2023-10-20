package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoOpenAccountIndexFind Open Account索引查询
// taobao.open.account.index.find
//
// Open Account索引查询
func TaobaoOpenAccountIndexFind(clt *core.SDKClient, req *user.TaobaoOpenAccountIndexFindAPIRequest, resp *user.TaobaoOpenAccountIndexFindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
