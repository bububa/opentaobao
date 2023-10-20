package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// Alibabaonetouchlogisticsexpresschargecalculate 计算快递运费&下单参数校验
// alibaba.onetouch.logistics.express.charge.calculate
//
// 计算快递运费、下单参数校验
func Alibabaonetouchlogisticsexpresschargecalculate(clt *core.SDKClient, req *icbulogistics.AlibabaonetouchlogisticsexpresschargecalculateAPIRequest, session string) (*icbulogistics.AlibabaonetouchlogisticsexpresschargecalculateAPIResponse, error) {
	var resp icbulogistics.AlibabaonetouchlogisticsexpresschargecalculateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
