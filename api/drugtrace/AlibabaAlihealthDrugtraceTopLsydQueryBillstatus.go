package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugtraceTopLsydQueryBillstatus
上传单据后处理状态查询
alibaba.alihealth.drugtrace.top.lsyd.query.billstatus

单据处理状态查询 */
func AlibabaAlihealthDrugtraceTopLsydQueryBillstatus(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
