package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytsearchbill 通过时间段批量查询入出库单信息
// alibaba.alihealth.drug.kyt.searchbill
//
// 通过时间段批量查询入出库单信息
func Alibabaalihealthdrugkytsearchbill(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytsearchbillAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytsearchbillAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytsearchbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
