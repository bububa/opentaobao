package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripTrainOrderSearch 获取火车票订单列表
// alitrip.btrip.train.order.search
//
// 第三方OA厂商获取自己的火车票数据
func AlitripBtripTrainOrderSearch(clt *core.SDKClient, req *btrip.AlitripBtripTrainOrderSearchAPIRequest, resp *btrip.AlitripBtripTrainOrderSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
