package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameJoincodeAssign 分配joinCode
// alibaba.cloudgame.interactive.game.joincode.assign
//
// 分配joinCode
func AlibabaCloudgameInteractiveGameJoincodeAssign(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest, resp *cloudgame.AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
