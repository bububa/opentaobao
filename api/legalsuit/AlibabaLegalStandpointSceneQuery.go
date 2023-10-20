package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalstandpointscenequery 查询场景
// alibaba.legal.standpoint.scene.query
//
// 查询场景
func Alibabalegalstandpointscenequery(clt *core.SDKClient, req *legalsuit.AlibabalegalstandpointscenequeryAPIRequest, session string) (*legalsuit.AlibabalegalstandpointscenequeryAPIResponse, error) {
	var resp legalsuit.AlibabalegalstandpointscenequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
