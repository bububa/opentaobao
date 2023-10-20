package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodekytydquerycode 查询追溯码对应的药品信息（药店）
// alibaba.alihealth.drug.code.kyt.yd.querycode
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
func Alibabaalihealthdrugcodekytydquerycode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodekytydquerycodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodekytydquerycodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodekytydquerycodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
