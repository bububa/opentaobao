package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaoopenaccountindexfind Open Account索引查询
// taobao.open.account.index.find
//
// Open Account索引查询
func Taobaoopenaccountindexfind(clt *core.SDKClient, req *user.TaobaoopenaccountindexfindAPIRequest, session string) (*user.TaobaoopenaccountindexfindAPIResponse, error) {
	var resp user.TaobaoopenaccountindexfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
