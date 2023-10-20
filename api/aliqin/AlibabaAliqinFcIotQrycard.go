package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcIotQrycard 查询终端信息
// alibaba.aliqin.fc.iot.qrycard
//
// 查询终端信息
func AlibabaAliqinFcIotQrycard(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotQrycardAPIRequest, resp *aliqin.AlibabaAliqinFcIotQrycardAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
