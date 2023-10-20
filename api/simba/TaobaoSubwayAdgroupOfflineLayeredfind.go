package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwayadgroupofflinelayeredfind 查询单元离线列表30天转化周期
// taobao.subway.adgroup.offline.layeredfind
//
// 查询单元离线列表
func Taobaosubwayadgroupofflinelayeredfind(clt *core.SDKClient, req *simba.TaobaosubwayadgroupofflinelayeredfindAPIRequest, session string) (*simba.TaobaosubwayadgroupofflinelayeredfindAPIResponse, error) {
	var resp simba.TaobaosubwayadgroupofflinelayeredfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
