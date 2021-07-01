package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugKytGetbyrefentid
根据企业唯一标识查看企业详细信息
alibaba.alihealth.drug.kyt.getbyrefentid

根据企业唯一标识查看企业详细信息 */
func AlibabaAlihealthDrugKytGetbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetbyrefentidAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytGetbyrefentidAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytGetbyrefentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
