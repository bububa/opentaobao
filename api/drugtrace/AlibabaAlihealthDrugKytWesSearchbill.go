package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesSearchbill 通过时间段批量查询入出库单信息
// alibaba.alihealth.drug.kyt.wes.searchbill
//
// 通过时间段批量查询入出库单信息
func AlibabaAlihealthDrugKytWesSearchbill(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesSearchbillAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytWesSearchbillAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytWesSearchbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
