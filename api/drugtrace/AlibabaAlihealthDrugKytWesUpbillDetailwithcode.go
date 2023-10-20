package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytwesupbilldetailwithcode 查询上游出库单明细（带追溯码信息）
// alibaba.alihealth.drug.kyt.wes.upbill.detailwithcode
//
// 查询上游出库单明细(带追溯码信息)
func Alibabaalihealthdrugkytwesupbilldetailwithcode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
