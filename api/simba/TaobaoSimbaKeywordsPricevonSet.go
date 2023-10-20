package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordspricevonset 设置一批关键词的信息
// taobao.simba.keywords.pricevon.set
//
// 设置一批关键词的信息，包含无线出价、计算机出价和关键词匹配方式
func Taobaosimbakeywordspricevonset(clt *core.SDKClient, req *simba.TaobaosimbakeywordspricevonsetAPIRequest, session string) (*simba.TaobaosimbakeywordspricevonsetAPIResponse, error) {
	var resp simba.TaobaosimbakeywordspricevonsetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
