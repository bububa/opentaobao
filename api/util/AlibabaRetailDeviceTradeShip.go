package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Alibabaretaildevicetradeship 贩卖机掉货成功通知
// alibaba.retail.device.trade.ship
//
// 贩卖机发货
func Alibabaretaildevicetradeship(clt *core.SDKClient, req *util.AlibabaretaildevicetradeshipAPIRequest, session string) (*util.AlibabaretaildevicetradeshipAPIResponse, error) {
	var resp util.AlibabaretaildevicetradeshipAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
