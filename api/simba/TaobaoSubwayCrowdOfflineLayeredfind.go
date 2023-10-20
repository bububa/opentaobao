package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwaycrowdofflinelayeredfind 获取人群离线报表30转化周期
// taobao.subway.crowd.offline.layeredfind
//
// 获取人群离线报表
func Taobaosubwaycrowdofflinelayeredfind(clt *core.SDKClient, req *simba.TaobaosubwaycrowdofflinelayeredfindAPIRequest, session string) (*simba.TaobaosubwaycrowdofflinelayeredfindAPIResponse, error) {
	var resp simba.TaobaosubwaycrowdofflinelayeredfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
