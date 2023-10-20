package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesQuerycodeactive 查询码是否激活
// alibaba.alihealth.drug.kyt.wes.querycodeactive
//
// 查询码是否激活
func AlibabaAlihealthDrugKytWesQuerycodeactive(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
