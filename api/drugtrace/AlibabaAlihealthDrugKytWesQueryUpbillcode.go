package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytwesqueryupbillcode 通过一个码查询上游出库单
// alibaba.alihealth.drug.kyt.wes.query.upbillcode
//
// 一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
func Alibabaalihealthdrugkytwesqueryupbillcode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytwesqueryupbillcodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytwesqueryupbillcodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytwesqueryupbillcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
