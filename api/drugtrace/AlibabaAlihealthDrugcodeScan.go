package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeScan 查询扫码信息
// alibaba.alihealth.drugcode.scan
//
// 查询扫码信息
func AlibabaAlihealthDrugcodeScan(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeScanAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeScanAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
