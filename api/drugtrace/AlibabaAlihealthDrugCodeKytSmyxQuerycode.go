package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodekytsmyxquerycode 扫码营销码查询
// alibaba.alihealth.drug.code.kyt.smyx.querycode
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
func Alibabaalihealthdrugcodekytsmyxquerycode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodekytsmyxquerycodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodekytsmyxquerycodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodekytsmyxquerycodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
