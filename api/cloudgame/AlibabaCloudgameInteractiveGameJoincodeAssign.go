package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacloudgameinteractivegamejoincodeassign 分配joinCode
// alibaba.cloudgame.interactive.game.joincode.assign
//
// 分配joinCode
func Alibabacloudgameinteractivegamejoincodeassign(clt *core.SDKClient, req *cloudgame.AlibabacloudgameinteractivegamejoincodeassignAPIRequest, session string) (*cloudgame.AlibabacloudgameinteractivegamejoincodeassignAPIResponse, error) {
	var resp cloudgame.AlibabacloudgameinteractivegamejoincodeassignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
