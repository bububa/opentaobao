package lstvending

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstvending"
)

// Alibabalstvendingorderupdate 自动售货机订单物流信息回传
// alibaba.lst.vending.order.update
//
// 零售通与设备供应商进行订单对接，通过此接口回流订单物流信息。
func Alibabalstvendingorderupdate(clt *core.SDKClient, req *lstvending.AlibabalstvendingorderupdateAPIRequest, session string) (*lstvending.AlibabalstvendingorderupdateAPIResponse, error) {
	var resp lstvending.AlibabalstvendingorderupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
