package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvPvQuery 查询pv属性
// alibaba.idle.isv.pv.query
//
// 查询pv属性
func AlibabaIdleIsvPvQuery(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvPvQueryAPIRequest, session string) (*idleisv.AlibabaIdleIsvPvQueryAPIResponse, error) {
	var resp idleisv.AlibabaIdleIsvPvQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
