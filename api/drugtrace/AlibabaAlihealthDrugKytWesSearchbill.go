package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytwessearchbill 通过时间段批量查询入出库单信息
// alibaba.alihealth.drug.kyt.wes.searchbill
//
// 通过时间段批量查询入出库单信息
func Alibabaalihealthdrugkytwessearchbill(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytwessearchbillAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytwessearchbillAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytwessearchbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
