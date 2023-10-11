package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcheckcodeChecklastfour 校验追溯码的后4位是否正确
// alibaba.alihealth.drugcheckcode.checklastfour
//
// 校验追溯码的后4位是否正确
func AlibabaAlihealthDrugcheckcodeChecklastfour(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcheckcodeChecklastfourAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
