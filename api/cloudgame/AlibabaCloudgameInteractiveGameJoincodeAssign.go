package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameJoincodeAssign 分配joinCode
// alibaba.cloudgame.interactive.game.joincode.assign
//
// 分配joinCode
func AlibabaCloudgameInteractiveGameJoincodeAssign(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameJoincodeAssignAPIRequest, session string) (*cloudgame.AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse, error) {
	var resp cloudgame.AlibabaCloudgameInteractiveGameJoincodeAssignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
