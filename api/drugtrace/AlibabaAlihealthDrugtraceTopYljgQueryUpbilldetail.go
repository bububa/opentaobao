package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetail 根据单据号查询单据的详情信息【注意：查询的是本企业的单据】
// alibaba.alihealth.drugtrace.top.yljg.query.upbilldetail
//
// 根据单据号查询单据的详情信息【注意：这个接口查询的是本企业的单据，如果是查询上游的单据明细信息，使用xxxxxxx.listupout.detail接口】。
func AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetail(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
