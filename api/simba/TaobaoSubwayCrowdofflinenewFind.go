package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayCrowdofflinenewFind 获取人群离线多日汇总报表
// taobao.subway.crowdofflinenew.find
//
// 获取人群离线报表
func TaobaoSubwayCrowdofflinenewFind(clt *core.SDKClient, req *simba.TaobaoSubwayCrowdofflinenewFindAPIRequest, session string) (*simba.TaobaoSubwayCrowdofflinenewFindAPIResponse, error) {
	var resp simba.TaobaoSubwayCrowdofflinenewFindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
