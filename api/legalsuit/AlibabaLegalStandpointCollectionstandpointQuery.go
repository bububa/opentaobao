package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalstandpointcollectionstandpointquery 查询收藏口径
// alibaba.legal.standpoint.collectionstandpoint.query
//
// 查询收藏口径
func Alibabalegalstandpointcollectionstandpointquery(clt *core.SDKClient, req *legalsuit.AlibabalegalstandpointcollectionstandpointqueryAPIRequest, session string) (*legalsuit.AlibabalegalstandpointcollectionstandpointqueryAPIResponse, error) {
	var resp legalsuit.AlibabalegalstandpointcollectionstandpointqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
