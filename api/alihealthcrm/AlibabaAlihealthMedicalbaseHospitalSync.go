package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthMedicalbaseHospitalSync 互联网医院批量导入接口
// alibaba.alihealth.medicalbase.hospital.sync
//
// 互联网医院isv批量通过接口批量导入
func AlibabaAlihealthMedicalbaseHospitalSync(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest, resp *alihealthcrm.AlibabaAlihealthMedicalbaseHospitalSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
