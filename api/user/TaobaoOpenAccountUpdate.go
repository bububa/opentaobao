package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaoopenaccountupdate Open Account数据更新
// taobao.open.account.update
//
// Open Account数据更新
func Taobaoopenaccountupdate(clt *core.SDKClient, req *user.TaobaoopenaccountupdateAPIRequest, session string) (*user.TaobaoopenaccountupdateAPIResponse, error) {
	var resp user.TaobaoopenaccountupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
