package shenjing

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shenjing"
)

// AlibabaIbShenjingVisitorPadOpendoor 访客发起开门
// alibaba.ib.shenjing.visitor.pad.opendoor
//
// 访客PAD端录入完人脸后，可以点击开门按钮开门。
func AlibabaIbShenjingVisitorPadOpendoor(clt *core.SDKClient, req *shenjing.AlibabaIbShenjingVisitorPadOpendoorAPIRequest, resp *shenjing.AlibabaIbShenjingVisitorPadOpendoorAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
