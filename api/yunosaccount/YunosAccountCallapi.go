package yunosaccount

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosaccount"
)

/* YunosAccountCallapi
调用YunOS账号开放API
yunos.account.callapi

YunOS账号客户端对外开放的api通过top暴露 */
func YunosAccountCallapi(clt *core.SDKClient, req *yunosaccount.YunosAccountCallapiAPIRequest, session string) (*yunosaccount.YunosAccountCallapiAPIResponse, error) {
	var resp yunosaccount.YunosAccountCallapiAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
