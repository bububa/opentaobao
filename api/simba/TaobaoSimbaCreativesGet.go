package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacreativesget 批量获得创意
// taobao.simba.creatives.get
//
// 取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；&lt;br/&gt;如果同时提供了推广组Id和创意id列表，则优先使用推广组Id；
func Taobaosimbacreativesget(clt *core.SDKClient, req *simba.TaobaosimbacreativesgetAPIRequest, session string) (*simba.TaobaosimbacreativesgetAPIResponse, error) {
	var resp simba.TaobaosimbacreativesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
