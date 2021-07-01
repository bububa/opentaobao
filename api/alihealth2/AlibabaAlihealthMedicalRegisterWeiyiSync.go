package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

/* AlibabaAlihealthMedicalRegisterWeiyiSync
微医数据号源回传
alibaba.alihealth.medical.register.weiyi.sync

微医号源数据回传 */
func AlibabaAlihealthMedicalRegisterWeiyiSync(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalRegisterWeiyiSyncAPIRequest, session string) (*alihealth2.AlibabaAlihealthMedicalRegisterWeiyiSyncAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthMedicalRegisterWeiyiSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
