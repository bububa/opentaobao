package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbasalestarcreativeadd （新）新建创意
// taobao.simba.salestar.creative.add
//
// 创建一个创意
func Taobaosimbasalestarcreativeadd(clt *core.SDKClient, req *simba.TaobaosimbasalestarcreativeaddAPIRequest, session string) (*simba.TaobaosimbasalestarcreativeaddAPIResponse, error) {
	var resp simba.TaobaosimbasalestarcreativeaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
