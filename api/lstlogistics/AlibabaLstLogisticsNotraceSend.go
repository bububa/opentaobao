package lstlogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// Alibabalstlogisticsnotracesend 供应商-异云-无需物流发货
// alibaba.lst.logistics.notrace.send
//
// 异地云仓的订单，使用无需物流的发货方式来变更订单发货状态
func Alibabalstlogisticsnotracesend(clt *core.SDKClient, req *lstlogistics.AlibabalstlogisticsnotracesendAPIRequest, session string) (*lstlogistics.AlibabalstlogisticsnotracesendAPIResponse, error) {
	var resp lstlogistics.AlibabalstlogisticsnotracesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
