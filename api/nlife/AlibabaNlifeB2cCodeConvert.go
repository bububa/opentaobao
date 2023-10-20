package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// Alibabanlifeb2ccodeconvert b2c转码
// alibaba.nlife.b2c.code.convert
//
// 将商品的URL转码，ISV将该码写入RFID
func Alibabanlifeb2ccodeconvert(clt *core.SDKClient, req *nlife.Alibabanlifeb2ccodeconvertAPIRequest, session string) (*nlife.Alibabanlifeb2ccodeconvertAPIResponse, error) {
	var resp nlife.Alibabanlifeb2ccodeconvertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
