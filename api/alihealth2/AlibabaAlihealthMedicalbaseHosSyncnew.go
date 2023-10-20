package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalbaseHosSyncnew 直连医院上传接口
// alibaba.alihealth.medicalbase.hos.syncnew
//
// 直连医院上传接口
func AlibabaAlihealthMedicalbaseHosSyncnew(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest, resp *alihealth2.AlibabaAlihealthMedicalbaseHosSyncnewAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
