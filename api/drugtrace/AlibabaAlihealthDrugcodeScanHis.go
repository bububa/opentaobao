package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeScanHis 企业查询扫码历史
// alibaba.alihealth.drugcode.scan.his
//
// 企业查询扫码历史
func AlibabaAlihealthDrugcodeScanHis(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeScanHisAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeScanHisAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
