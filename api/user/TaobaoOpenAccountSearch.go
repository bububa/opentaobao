package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaoopenaccountsearch open account数据搜索
// taobao.open.account.search
//
// open account数据搜索
func Taobaoopenaccountsearch(clt *core.SDKClient, req *user.TaobaoopenaccountsearchAPIRequest, session string) (*user.TaobaoopenaccountsearchAPIResponse, error) {
	var resp user.TaobaoopenaccountsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
