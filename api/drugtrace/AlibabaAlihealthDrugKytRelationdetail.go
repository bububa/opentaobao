package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytRelationdetail 关联关系处理详情
// alibaba.alihealth.drug.kyt.relationdetail
//
// 关联关系处理详情
func AlibabaAlihealthDrugKytRelationdetail(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytRelationdetailAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytRelationdetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
