package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalstandpointstandpointquery 查询具体口径
// alibaba.legal.standpoint.standpoint.query
//
// 查询具体口径
func Alibabalegalstandpointstandpointquery(clt *core.SDKClient, req *legalsuit.AlibabalegalstandpointstandpointqueryAPIRequest, session string) (*legalsuit.AlibabalegalstandpointstandpointqueryAPIResponse, error) {
	var resp legalsuit.AlibabalegalstandpointstandpointqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
