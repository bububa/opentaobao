package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwayciaupdate 批量修改单元智能出价
// taobao.subway.cia.update
//
// 批量修改直通车推广单元的智能出价配置
func Taobaosubwayciaupdate(clt *core.SDKClient, req *simba.TaobaosubwayciaupdateAPIRequest, session string) (*simba.TaobaosubwayciaupdateAPIResponse, error) {
	var resp simba.TaobaosubwayciaupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
