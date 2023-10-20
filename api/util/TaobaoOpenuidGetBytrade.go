package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaoopenuidgetbytrade 通过订单获取对应买家的openUID
// taobao.openuid.get.bytrade
//
// 通过订单获取对应买家的openUID,需要卖家授权
func Taobaoopenuidgetbytrade(clt *core.SDKClient, req *util.TaobaoopenuidgetbytradeAPIRequest, session string) (*util.TaobaoopenuidgetbytradeAPIResponse, error) {
	var resp util.TaobaoopenuidgetbytradeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
