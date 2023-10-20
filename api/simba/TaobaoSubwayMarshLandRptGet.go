package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwaymarshlandrptget 获取捡漏词包分时报表数据
// taobao.subway.marsh.land.rpt.get
//
// 获取捡漏词包分时报表数据
func Taobaosubwaymarshlandrptget(clt *core.SDKClient, req *simba.TaobaosubwaymarshlandrptgetAPIRequest, session string) (*simba.TaobaosubwaymarshlandrptgetAPIResponse, error) {
	var resp simba.TaobaosubwaymarshlandrptgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
