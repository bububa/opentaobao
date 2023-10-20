package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthtracecodesellerchannelsearch 查询渠道商api
// alibaba.alihealth.tracecodeseller.channel.search
//
// 查询渠道商api
func Alibabaalihealthtracecodesellerchannelsearch(clt *core.SDKClient, req *alihealth2.AlibabaalihealthtracecodesellerchannelsearchAPIRequest, session string) (*alihealth2.AlibabaalihealthtracecodesellerchannelsearchAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthtracecodesellerchannelsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
