package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaoopenaccountdelete OpenAccount删除数据
// taobao.open.account.delete
//
// OpenAccount删除数据
func Taobaoopenaccountdelete(clt *core.SDKClient, req *user.TaobaoopenaccountdeleteAPIRequest, session string) (*user.TaobaoopenaccountdeleteAPIResponse, error) {
	var resp user.TaobaoopenaccountdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
