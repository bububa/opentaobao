package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// AlibabaNlifeB2cTradestatusDrive b2c订单状态驱动
// alibaba.nlife.b2c.tradestatus.drive
//
// 用于驱动零售+订单状态
func AlibabaNlifeB2cTradestatusDrive(clt *core.SDKClient, req *nlife.AlibabaNlifeB2cTradestatusDriveAPIRequest, session string) (*nlife.AlibabaNlifeB2cTradestatusDriveAPIResponse, error) {
	var resp nlife.AlibabaNlifeB2cTradestatusDriveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
