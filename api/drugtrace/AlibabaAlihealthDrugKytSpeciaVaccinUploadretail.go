package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSpeciaVaccinUploadretail 零售单据上传接口（疫苗）
// alibaba.alihealth.drug.kyt.specia.vaccin.uploadretail
//
// 零售上传单据信息接口
func AlibabaAlihealthDrugKytSpeciaVaccinUploadretail(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
