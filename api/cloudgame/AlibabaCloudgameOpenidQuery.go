package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacloudgameopenidquery 咖哒用户ID查询
// alibaba.cloudgame.openid.query
//
// 云游戏业务需要提供查询用户信息的TOP能力
func Alibabacloudgameopenidquery(clt *core.SDKClient, req *cloudgame.AlibabacloudgameopenidqueryAPIRequest, session string) (*cloudgame.AlibabacloudgameopenidqueryAPIResponse, error) {
	var resp cloudgame.AlibabacloudgameopenidqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
