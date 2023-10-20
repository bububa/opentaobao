package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkreversetimeslice 逆向取货时间片查询
// alibaba.wdk.reverse.timeslice
//
// 逆向取货时间片查询
func Alibabawdkreversetimeslice(clt *core.SDKClient, req *wdk.AlibabawdkreversetimesliceAPIRequest, session string) (*wdk.AlibabawdkreversetimesliceAPIResponse, error) {
	var resp wdk.AlibabawdkreversetimesliceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
