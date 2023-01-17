package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesSearchbillDetail 查询单据详情
// alibaba.alihealth.drug.kyt.wes.searchbill.detail
//
// 根据单据号码查询码单据详情和码信息
func AlibabaAlihealthDrugKytWesSearchbillDetail(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesSearchbillDetailAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytWesSearchbillDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
