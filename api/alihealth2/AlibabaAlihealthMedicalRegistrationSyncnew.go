package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

/* AlibabaAlihealthMedicalRegistrationSyncnew
阿里健康新挂号数据回传
alibaba.alihealth.medical.registration.syncnew

阿里健康新挂号记录回传接口 */
func AlibabaAlihealthMedicalRegistrationSyncnew(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest, session string) (*alihealth2.AlibabaAlihealthMedicalRegistrationSyncnewAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthMedicalRegistrationSyncnewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
