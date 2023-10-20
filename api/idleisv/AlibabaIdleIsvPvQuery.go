package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvPvQuery 查询pv属性
// alibaba.idle.isv.pv.query
//
// 查询pv属性
func AlibabaIdleIsvPvQuery(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvPvQueryAPIRequest, resp *idleisv.AlibabaIdleIsvPvQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
