package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentid 根据企业唯一标识查看企业详细信息
// alibaba.alihealth.drugtrace.top.lsyd.query.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
func AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
