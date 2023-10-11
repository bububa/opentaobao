package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesGetbyentid 根据企业主键查看企业详细信息
// alibaba.alihealth.drug.kyt.wes.getbyentid
//
// 根据企业主键查看企业详细信息
func AlibabaAlihealthDrugKytWesGetbyentid(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesGetbyentidAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytWesGetbyentidAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytWesGetbyentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
