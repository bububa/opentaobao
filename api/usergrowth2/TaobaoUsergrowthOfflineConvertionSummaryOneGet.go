package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// TaobaoUsergrowthOfflineConvertionSummaryOneGet 手淘用增 线下业务 转化t+1汇总数据
// taobao.usergrowth.offline.convertion.summary.one.get
//
// 淘系用户增长团队-线下拉新业务，对线下渠道提供mapi，目的是为了给有开发能力的渠道提供返数功能，方便渠道对手下的推广者结算（t+1汇总）
func TaobaoUsergrowthOfflineConvertionSummaryOneGet(clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthOfflineConvertionSummaryOneGetAPIRequest, session string) (*usergrowth2.TaobaoUsergrowthOfflineConvertionSummaryOneGetAPIResponse, error) {
	var resp usergrowth2.TaobaoUsergrowthOfflineConvertionSummaryOneGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
