package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanItemSubscribeDailyLeftQuery 查询当天可添加的余量
// taobao.baichuan.item.subscribe.daily.left.query
//
// 查询当天可添加的余量
func TaobaoBaichuanItemSubscribeDailyLeftQuery(clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemSubscribeDailyLeftQueryAPIRequest, resp *baichuan.TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
