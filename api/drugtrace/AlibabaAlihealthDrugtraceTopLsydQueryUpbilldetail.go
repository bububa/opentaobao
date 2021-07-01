package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetail
根据单据号查询单据的详情信息【注意：查询的是本企业的单据】
alibaba.alihealth.drugtrace.top.lsyd.query.upbilldetail

根据单据号查询单据的详情信息【注意：这个接口查询的是本企业的单据，如果是查询上游的单据明细信息，使用xxxxxxx.listupout.detail接口】。 */
func AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetail(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
