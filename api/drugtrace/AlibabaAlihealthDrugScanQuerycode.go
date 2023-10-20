package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugscanquerycode 查询药监码对应的有效期和包装规格
// alibaba.alihealth.drug.scan.querycode
//
// 查询药监码对应的有效期和包装规格
func Alibabaalihealthdrugscanquerycode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugscanquerycodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugscanquerycodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugscanquerycodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
