package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytYyQuerysubcodes 查询一个码的所有子码
// alibaba.alihealth.drug.kyt.yy.querysubcodes
//
// 单码的了码查询
func AlibabaAlihealthDrugKytYyQuerysubcodes(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytYyQuerysubcodesAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytYyQuerysubcodesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
