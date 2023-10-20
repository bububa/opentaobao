package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwayadgroupofflinefind 查询单元离线多日汇总列表
// taobao.subway.adgroup.offline.find
//
// 查询单元离线列表
func Taobaosubwayadgroupofflinefind(clt *core.SDKClient, req *simba.TaobaosubwayadgroupofflinefindAPIRequest, session string) (*simba.TaobaosubwayadgroupofflinefindAPIResponse, error) {
	var resp simba.TaobaosubwayadgroupofflinefindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
