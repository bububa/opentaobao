package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcIotCardStatus 物联卡状态查询
// alibaba.aliqin.fc.iot.cardStatus
//
// 物联卡状态查询
func AlibabaAliqinFcIotCardStatus(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotCardStatusAPIRequest, resp *aliqin.AlibabaAliqinFcIotCardStatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
