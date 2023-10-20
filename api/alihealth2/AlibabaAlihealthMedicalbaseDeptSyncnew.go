package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalbaseDeptSyncnew 直连科室上传
// alibaba.alihealth.medicalbase.dept.syncnew
//
// 直连科室上传接口
func AlibabaAlihealthMedicalbaseDeptSyncnew(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalbaseDeptSyncnewAPIRequest, session string) (*alihealth2.AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
