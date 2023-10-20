package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthtracecodesellerentsearch 查询商家信息
// alibaba.alihealth.tracecodeseller.ent.search
//
// 查询商家信息
func Alibabaalihealthtracecodesellerentsearch(clt *core.SDKClient, req *alihealth2.AlibabaalihealthtracecodesellerentsearchAPIRequest, session string) (*alihealth2.AlibabaalihealthtracecodesellerentsearchAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthtracecodesellerentsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
