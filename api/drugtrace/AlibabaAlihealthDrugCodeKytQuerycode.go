package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodekytquerycode 查询追溯码对应的药品信息
// alibaba.alihealth.drug.code.kyt.querycode
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
func Alibabaalihealthdrugcodekytquerycode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodekytquerycodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodekytquerycodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodekytquerycodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
