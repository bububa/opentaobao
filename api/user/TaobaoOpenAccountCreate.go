package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaoopenaccountcreate Open Account导入数据
// taobao.open.account.create
//
// Open Account导入数据
func Taobaoopenaccountcreate(clt *core.SDKClient, req *user.TaobaoopenaccountcreateAPIRequest, session string) (*user.TaobaoopenaccountcreateAPIResponse, error) {
	var resp user.TaobaoopenaccountcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
