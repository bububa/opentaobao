package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugDownloadDataerrordiagnosis 数据未落地诊断
// alibaba.alihealth.drug.download.dataerrordiagnosis
//
// 阿里健康-追溯码-D2D数据未落地原因诊断
func AlibabaAlihealthDrugDownloadDataerrordiagnosis(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest, resp *drugtrace.AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
