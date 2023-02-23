package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugDownloadDataerrordiagnosis 数据未落地诊断
// alibaba.alihealth.drug.download.dataerrordiagnosis
//
// 阿里健康-追溯码-D2D数据未落地原因诊断
func AlibabaAlihealthDrugDownloadDataerrordiagnosis(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
