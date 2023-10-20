package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeListCodeGov 政府码查询接口
// alibaba.alihealth.drug.code.list.code.gov
//
// 政府码查询接口
func AlibabaAlihealthDrugCodeListCodeGov(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeListCodeGovAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugCodeListCodeGovAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugCodeListCodeGovAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
