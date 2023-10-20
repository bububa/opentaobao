package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytquerycoderelationfrombillcode 根据单据号码查询码单据详情和码信息
// alibaba.alihealth.drug.kyt.query.code.relation.from.billcode
//
// 根据单据号码查询码单据详情和码信息
func Alibabaalihealthdrugkytquerycoderelationfrombillcode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytquerycoderelationfrombillcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
