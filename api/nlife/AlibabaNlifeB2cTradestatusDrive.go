package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// Alibabanlifeb2ctradestatusdrive b2c订单状态驱动
// alibaba.nlife.b2c.tradestatus.drive
//
// 用于驱动零售+订单状态
func Alibabanlifeb2ctradestatusdrive(clt *core.SDKClient, req *nlife.Alibabanlifeb2ctradestatusdriveAPIRequest, session string) (*nlife.Alibabanlifeb2ctradestatusdriveAPIResponse, error) {
	var resp nlife.Alibabanlifeb2ctradestatusdriveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
