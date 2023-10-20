package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwayciaget 查询单元智能出价信息
// taobao.subway.cia.get
//
// 查询单元智能出价信息
func Taobaosubwayciaget(clt *core.SDKClient, req *simba.TaobaosubwayciagetAPIRequest, session string) (*simba.TaobaosubwayciagetAPIResponse, error) {
	var resp simba.TaobaosubwayciagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
