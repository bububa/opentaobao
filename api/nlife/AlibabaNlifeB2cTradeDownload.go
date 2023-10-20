package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// Alibabanlifeb2ctradedownload b2c下载订单
// alibaba.nlife.b2c.trade.download
//
// 下载零售商在零售+平台创建的订单
func Alibabanlifeb2ctradedownload(clt *core.SDKClient, req *nlife.Alibabanlifeb2ctradedownloadAPIRequest, session string) (*nlife.Alibabanlifeb2ctradedownloadAPIResponse, error) {
	var resp nlife.Alibabanlifeb2ctradedownloadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
