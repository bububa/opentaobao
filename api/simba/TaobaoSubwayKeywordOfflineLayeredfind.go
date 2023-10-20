package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwaykeywordofflinelayeredfind 查询关键词离线报表30天转化周期
// taobao.subway.keyword.offline.layeredfind
//
// 获取关键词离线报表
func Taobaosubwaykeywordofflinelayeredfind(clt *core.SDKClient, req *simba.TaobaosubwaykeywordofflinelayeredfindAPIRequest, session string) (*simba.TaobaosubwaykeywordofflinelayeredfindAPIResponse, error) {
	var resp simba.TaobaosubwaykeywordofflinelayeredfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
