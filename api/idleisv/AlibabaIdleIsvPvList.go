package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvPvList 闲鱼已验货pv查询
// alibaba.idle.isv.pv.list
//
// 根据闲鱼渠道类目查询对应的品牌和型号清单，供服务商进行选择
func AlibabaIdleIsvPvList(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvPvListAPIRequest, session string) (*idleisv.AlibabaIdleIsvPvListAPIResponse, error) {
	var resp idleisv.AlibabaIdleIsvPvListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
