package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugbillupbilldetailwithcode 查询上游出库单明细(带追溯码信息)
// alibaba.alihealth.drug.bill.upbill.detail.withcode
//
// 查询上游出库单明细(带追溯码信息)
func Alibabaalihealthdrugbillupbilldetailwithcode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugbillupbilldetailwithcodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugbillupbilldetailwithcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
