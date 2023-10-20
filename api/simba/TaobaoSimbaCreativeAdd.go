package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacreativeadd 增加创意
// taobao.simba.creative.add
//
// 创建一个创意
func Taobaosimbacreativeadd(clt *core.SDKClient, req *simba.TaobaosimbacreativeaddAPIRequest, session string) (*simba.TaobaosimbacreativeaddAPIResponse, error) {
	var resp simba.TaobaosimbacreativeaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
