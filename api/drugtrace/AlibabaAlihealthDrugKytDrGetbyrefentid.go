package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugKytDrGetbyrefentid
多融通过一个企业唯一标识查询企业详细信息
alibaba.alihealth.drug.kyt.dr.getbyrefentid

根据企业唯一标识查看企业详细信息 */
func AlibabaAlihealthDrugKytDrGetbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
