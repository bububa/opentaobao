package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwaycreativeofflinefind 获取创意离线多日汇总报表
// taobao.subway.creativeoffline.find
//
// 获取创意离线报表
func Taobaosubwaycreativeofflinefind(clt *core.SDKClient, req *simba.TaobaosubwaycreativeofflinefindAPIRequest, session string) (*simba.TaobaosubwaycreativeofflinefindAPIResponse, error) {
	var resp simba.TaobaosubwaycreativeofflinefindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
