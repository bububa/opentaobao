package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwaycreativeofflinelayeredfind 获取创意离线报表周期30天
// taobao.subway.creative.offline.layeredfind
//
// 获取创意离线报表
func Taobaosubwaycreativeofflinelayeredfind(clt *core.SDKClient, req *simba.TaobaosubwaycreativeofflinelayeredfindAPIRequest, session string) (*simba.TaobaosubwaycreativeofflinelayeredfindAPIResponse, error) {
	var resp simba.TaobaosubwaycreativeofflinelayeredfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
