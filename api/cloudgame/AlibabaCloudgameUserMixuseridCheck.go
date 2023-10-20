package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacloudgameusermixuseridcheck 云游戏混淆用户ID校验
// alibaba.cloudgame.user.mixuserid.check
//
// 验证混淆用户ID是否合法
func Alibabacloudgameusermixuseridcheck(clt *core.SDKClient, req *cloudgame.AlibabacloudgameusermixuseridcheckAPIRequest, session string) (*cloudgame.AlibabacloudgameusermixuseridcheckAPIResponse, error) {
	var resp cloudgame.AlibabacloudgameusermixuseridcheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
