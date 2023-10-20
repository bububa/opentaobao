package blackvip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/blackvip"
)

// Taobaoblackvipuserinfoget 88VIP用户信息查询
// taobao.blackvip.userinfo.get
//
// 查询88VIP用户信息，比如用户是否是88VIP，88VIP的失效时间等
func Taobaoblackvipuserinfoget(clt *core.SDKClient, req *blackvip.TaobaoblackvipuserinfogetAPIRequest, session string) (*blackvip.TaobaoblackvipuserinfogetAPIResponse, error) {
	var resp blackvip.TaobaoblackvipuserinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
