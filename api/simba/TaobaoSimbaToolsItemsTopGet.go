package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbatoolsitemstopget 取得一个关键词的推广组排名列表
// taobao.simba.tools.items.top.get
//
// 取得一个关键词的推广组排名列表
func Taobaosimbatoolsitemstopget(clt *core.SDKClient, req *simba.TaobaosimbatoolsitemstopgetAPIRequest, session string) (*simba.TaobaosimbatoolsitemstopgetAPIResponse, error) {
	var resp simba.TaobaosimbatoolsitemstopgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
