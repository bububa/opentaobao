package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacreativedelete 删除创意
// taobao.simba.creative.delete
//
// 删除一个创意
func Taobaosimbacreativedelete(clt *core.SDKClient, req *simba.TaobaosimbacreativedeleteAPIRequest, session string) (*simba.TaobaosimbacreativedeleteAPIResponse, error) {
	var resp simba.TaobaosimbacreativedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
