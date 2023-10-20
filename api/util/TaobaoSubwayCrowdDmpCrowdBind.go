package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaosubwaycrowddmpcrowdbind 直通车绑定达摩盘人群
// taobao.subway.crowd.dmp.crowd.bind
//
// 直通车绑定达摩盘人群
func Taobaosubwaycrowddmpcrowdbind(clt *core.SDKClient, req *util.TaobaosubwaycrowddmpcrowdbindAPIRequest, session string) (*util.TaobaosubwaycrowddmpcrowdbindAPIResponse, error) {
	var resp util.TaobaosubwaycrowddmpcrowdbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
