package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleconsignmentspustatistics 闲鱼帮卖同步服务商交易统计信息
// alibaba.idle.consignment.spu.statistics
//
// 闲鱼帮卖同步服务商交易统计信息
func Alibabaidleconsignmentspustatistics(clt *core.SDKClient, req *idle.AlibabaidleconsignmentspustatisticsAPIRequest, session string) (*idle.AlibabaidleconsignmentspustatisticsAPIResponse, error) {
	var resp idle.AlibabaidleconsignmentspustatisticsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
