package auction

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/auction"
)

// Taobaoauctiongovdatarealtimeget 获取实时(今日)统计数据
// taobao.auction.gov.data.realtime.get
//
// 提供查询当日法院及下属法院的拍卖统计数据
func Taobaoauctiongovdatarealtimeget(clt *core.SDKClient, req *auction.TaobaoauctiongovdatarealtimegetAPIRequest, session string) (*auction.TaobaoauctiongovdatarealtimegetAPIResponse, error) {
	var resp auction.TaobaoauctiongovdatarealtimegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
