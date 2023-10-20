package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwaycrowdofflinenewfind 获取人群离线多日汇总报表
// taobao.subway.crowdofflinenew.find
//
// 获取人群离线报表
func Taobaosubwaycrowdofflinenewfind(clt *core.SDKClient, req *simba.TaobaosubwaycrowdofflinenewfindAPIRequest, session string) (*simba.TaobaosubwaycrowdofflinenewfindAPIResponse, error) {
	var resp simba.TaobaosubwaycrowdofflinenewfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
