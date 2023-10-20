package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayWordpackageUpdate 批量更新词包
// taobao.subway.wordpackage.update
//
// 批量更新词包
func TaobaoSubwayWordpackageUpdate(clt *core.SDKClient, req *simba.TaobaoSubwayWordpackageUpdateAPIRequest, resp *simba.TaobaoSubwayWordpackageUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
