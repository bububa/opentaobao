package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesGetbyentid 根据企业主键查看企业详细信息
// alibaba.alihealth.drug.kyt.wes.getbyentid
//
// 根据企业主键查看企业详细信息
func AlibabaAlihealthDrugKytWesGetbyentid(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesGetbyentidAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesGetbyentidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
