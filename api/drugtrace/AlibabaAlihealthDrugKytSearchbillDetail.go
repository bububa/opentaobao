package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytsearchbilldetail 查询单据详情
// alibaba.alihealth.drug.kyt.searchbill.detail
//
// 根据单据号码查询码单据详情和码信息
func Alibabaalihealthdrugkytsearchbilldetail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytsearchbilldetailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytsearchbilldetailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytsearchbilldetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
