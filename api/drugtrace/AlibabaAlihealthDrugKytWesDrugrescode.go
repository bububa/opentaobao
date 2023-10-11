package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesDrugrescode 查询药品码段信息
// alibaba.alihealth.drug.kyt.wes.drugrescode
//
// 查询药品码段信息
func AlibabaAlihealthDrugKytWesDrugrescode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytWesDrugrescodeAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytWesDrugrescodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
