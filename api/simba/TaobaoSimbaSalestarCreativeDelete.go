package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbasalestarcreativedelete (新)销量明星删除创意相关接口
// taobao.simba.salestar.creative.delete
//
// 删除一个创意
func Taobaosimbasalestarcreativedelete(clt *core.SDKClient, req *simba.TaobaosimbasalestarcreativedeleteAPIRequest, session string) (*simba.TaobaosimbasalestarcreativedeleteAPIResponse, error) {
	var resp simba.TaobaosimbasalestarcreativedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
