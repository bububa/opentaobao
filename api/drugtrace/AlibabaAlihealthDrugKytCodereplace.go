package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytCodereplace 单码替换
// alibaba.alihealth.drug.kyt.codereplace
//
// 单码替换
func AlibabaAlihealthDrugKytCodereplace(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytCodereplaceAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytCodereplaceAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytCodereplaceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
