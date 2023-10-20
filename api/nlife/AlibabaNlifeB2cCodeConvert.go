package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// AlibabaNlifeB2cCodeConvert b2c转码
// alibaba.nlife.b2c.code.convert
//
// 将商品的URL转码，ISV将该码写入RFID
func AlibabaNlifeB2cCodeConvert(clt *core.SDKClient, req *nlife.AlibabaNlifeB2cCodeConvertAPIRequest, session string) (*nlife.AlibabaNlifeB2cCodeConvertAPIResponse, error) {
	var resp nlife.AlibabaNlifeB2cCodeConvertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
