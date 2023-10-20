package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytwessearchbilldetail 查询单据详情
// alibaba.alihealth.drug.kyt.wes.searchbill.detail
//
// 根据单据号码查询码单据详情和码信息
func Alibabaalihealthdrugkytwessearchbilldetail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytwessearchbilldetailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytwessearchbilldetailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytwessearchbilldetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
