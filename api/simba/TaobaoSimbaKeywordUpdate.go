package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordupdate （新）关键词更新相关接口
// taobao.simba.keyword.update
//
// （新）关键词更新相关接口
func Taobaosimbakeywordupdate(clt *core.SDKClient, req *simba.TaobaosimbakeywordupdateAPIRequest, session string) (*simba.TaobaosimbakeywordupdateAPIResponse, error) {
	var resp simba.TaobaosimbakeywordupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
