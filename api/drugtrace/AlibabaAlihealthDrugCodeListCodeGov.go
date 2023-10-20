package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeListCodeGov 政府码查询接口
// alibaba.alihealth.drug.code.list.code.gov
//
// 政府码查询接口
func AlibabaAlihealthDrugCodeListCodeGov(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeListCodeGovAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeListCodeGovAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
