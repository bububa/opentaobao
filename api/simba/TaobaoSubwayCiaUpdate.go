package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

/* TaobaoSubwayCiaUpdate
批量修改单元智能出价
taobao.subway.cia.update

批量修改直通车推广单元的智能出价配置 */
func TaobaoSubwayCiaUpdate(clt *core.SDKClient, req *simba.TaobaoSubwayCiaUpdateAPIRequest, session string) (*simba.TaobaoSubwayCiaUpdateAPIResponse, error) {
	var resp simba.TaobaoSubwayCiaUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
