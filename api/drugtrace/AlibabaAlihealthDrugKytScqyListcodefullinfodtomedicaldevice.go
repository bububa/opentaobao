package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldevice 医疗器械的码查询信息接口
// alibaba.alihealth.drug.kyt.scqy.listcodefullinfodtomedicaldevice
//
// 医疗器械的码查询信息接口
func AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldevice(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
