package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcIotCardoffer 查询物联网卡上订购的offer
// alibaba.aliqin.fc.iot.cardoffer
//
// 查询物联网卡上订购的offer
func AlibabaAliqinFcIotCardoffer(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotCardofferAPIRequest, resp *aliqin.AlibabaAliqinFcIotCardofferAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
