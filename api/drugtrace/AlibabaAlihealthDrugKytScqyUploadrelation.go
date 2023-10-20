package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqyUploadrelation 关联关系上传
// alibaba.alihealth.drug.kyt.scqy.uploadrelation
//
// 关联关系上传
func AlibabaAlihealthDrugKytScqyUploadrelation(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytScqyUploadrelationAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytScqyUploadrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
