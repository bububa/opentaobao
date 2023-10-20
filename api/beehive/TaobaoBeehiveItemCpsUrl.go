package beehive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/beehive"
)

// Taobaobeehiveitemcpsurl 分佣链接生成接口
// taobao.beehive.item.cps.url
//
// 传入包括itemId,accountId,bizType在内的参数，对应参数返回分佣链接
func Taobaobeehiveitemcpsurl(clt *core.SDKClient, req *beehive.TaobaobeehiveitemcpsurlAPIRequest, session string) (*beehive.TaobaobeehiveitemcpsurlAPIResponse, error) {
	var resp beehive.TaobaobeehiveitemcpsurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
