package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodesellerChannelSearch 查询渠道商api
// alibaba.alihealth.tracecodeseller.channel.search
//
// 查询渠道商api
func AlibabaAlihealthTracecodesellerChannelSearch(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerChannelSearchAPIRequest, resp *alihealth2.AlibabaAlihealthTracecodesellerChannelSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
