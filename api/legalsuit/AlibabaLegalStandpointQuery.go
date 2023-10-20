package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalstandpointquery 口径查询
// alibaba.legal.standpoint.query
//
// 口径查询
func Alibabalegalstandpointquery(clt *core.SDKClient, req *legalsuit.AlibabalegalstandpointqueryAPIRequest, session string) (*legalsuit.AlibabalegalstandpointqueryAPIResponse, error) {
	var resp legalsuit.AlibabalegalstandpointqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
