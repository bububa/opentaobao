package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeScanHis 企业查询扫码历史
// alibaba.alihealth.drugcode.scan.his
//
// 企业查询扫码历史
func AlibabaAlihealthDrugcodeScanHis(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeScanHisAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeScanHisAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
