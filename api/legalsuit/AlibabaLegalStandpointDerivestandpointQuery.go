package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalstandpointderivestandpointquery 查询衍生口径
// alibaba.legal.standpoint.derivestandpoint.query
//
// 查询衍生口径
func Alibabalegalstandpointderivestandpointquery(clt *core.SDKClient, req *legalsuit.AlibabalegalstandpointderivestandpointqueryAPIRequest, session string) (*legalsuit.AlibabalegalstandpointderivestandpointqueryAPIResponse, error) {
	var resp legalsuit.AlibabalegalstandpointderivestandpointqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
