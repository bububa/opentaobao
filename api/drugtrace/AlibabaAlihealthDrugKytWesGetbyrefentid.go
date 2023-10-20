package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesGetbyrefentid 根据企业唯一标识查看企业详细信息
// alibaba.alihealth.drug.kyt.wes.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
func AlibabaAlihealthDrugKytWesGetbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesGetbyrefentidAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
