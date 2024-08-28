package nlife

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// AlibabaNlifeB2cCodeConvert b2c转码
// alibaba.nlife.b2c.code.convert
//
// 将商品的URL转码，ISV将该码写入RFID
func AlibabaNlifeB2cCodeConvert(ctx context.Context, clt *core.SDKClient, req *nlife.AlibabaNlifeB2cCodeConvertAPIRequest, resp *nlife.AlibabaNlifeB2cCodeConvertAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
