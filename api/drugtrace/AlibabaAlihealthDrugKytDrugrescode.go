package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugKytDrugrescode
查询药品码段信息
alibaba.alihealth.drug.kyt.drugrescode

查询药品码段信息 */
func AlibabaAlihealthDrugKytDrugrescode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrugrescodeAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrugrescodeAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytDrugrescodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
