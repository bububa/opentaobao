package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalstandpointstandpointqueryall 滑动查询口径
// alibaba.legal.standpoint.standpoint.queryall
//
// 滑动查询口径
func Alibabalegalstandpointstandpointqueryall(clt *core.SDKClient, req *legalsuit.AlibabalegalstandpointstandpointqueryallAPIRequest, session string) (*legalsuit.AlibabalegalstandpointstandpointqueryallAPIResponse, error) {
	var resp legalsuit.AlibabalegalstandpointstandpointqueryallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
