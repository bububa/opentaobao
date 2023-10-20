package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalstandpointstandpointtreequery 查询口径树目录
// alibaba.legal.standpoint.standpointtree.query
//
// 查询口径树目录
func Alibabalegalstandpointstandpointtreequery(clt *core.SDKClient, req *legalsuit.AlibabalegalstandpointstandpointtreequeryAPIRequest, session string) (*legalsuit.AlibabalegalstandpointstandpointtreequeryAPIResponse, error) {
	var resp legalsuit.AlibabalegalstandpointstandpointtreequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
