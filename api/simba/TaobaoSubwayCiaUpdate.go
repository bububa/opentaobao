package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayCiaUpdate 批量修改单元智能出价
// taobao.subway.cia.update
//
// 批量修改直通车推广单元的智能出价配置
func TaobaoSubwayCiaUpdate(clt *core.SDKClient, req *simba.TaobaoSubwayCiaUpdateAPIRequest, resp *simba.TaobaoSubwayCiaUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
