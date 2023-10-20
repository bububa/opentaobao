package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordadd （新）关键词新增接口
// taobao.simba.keyword.add
//
// （新）关键词更新相关接口
func Taobaosimbakeywordadd(clt *core.SDKClient, req *simba.TaobaosimbakeywordaddAPIRequest, session string) (*simba.TaobaosimbakeywordaddAPIResponse, error) {
	var resp simba.TaobaosimbakeywordaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
