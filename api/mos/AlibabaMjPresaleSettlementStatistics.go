package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjpresalesettlementstatistics 预购结算数据统计
// alibaba.mj.presale.settlement.statistics
//
// 预购结算数据统计
func Alibabamjpresalesettlementstatistics(clt *core.SDKClient, req *mos.AlibabamjpresalesettlementstatisticsAPIRequest, session string) (*mos.AlibabamjpresalesettlementstatisticsAPIResponse, error) {
	var resp mos.AlibabamjpresalesettlementstatisticsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
