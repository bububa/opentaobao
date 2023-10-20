package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalstandpointstandpointcollection 收藏|取消收藏
// alibaba.legal.standpoint.standpoint.collection
//
// 收藏|取消收藏
func Alibabalegalstandpointstandpointcollection(clt *core.SDKClient, req *legalsuit.AlibabalegalstandpointstandpointcollectionAPIRequest, session string) (*legalsuit.AlibabalegalstandpointstandpointcollectionAPIResponse, error) {
	var resp legalsuit.AlibabalegalstandpointstandpointcollectionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
