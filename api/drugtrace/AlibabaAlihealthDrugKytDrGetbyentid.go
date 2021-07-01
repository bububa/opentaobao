package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugKytDrGetbyentid
多融通过企业ID得到一个企业的详细信息
alibaba.alihealth.drug.kyt.dr.getbyentid

根据企业主键查看企业详细信息 */
func AlibabaAlihealthDrugKytDrGetbyentid(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrGetbyentidAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrGetbyentidAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytDrGetbyentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
