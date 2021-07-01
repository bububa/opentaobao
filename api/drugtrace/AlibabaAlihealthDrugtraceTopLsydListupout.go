package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugtraceTopLsydListupout
零售药店查询本企业上游企业出库单据信息
alibaba.alihealth.drugtrace.top.lsyd.listupout

查询货主/本企业上游企业出库单据信息 */
func AlibabaAlihealthDrugtraceTopLsydListupout(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
