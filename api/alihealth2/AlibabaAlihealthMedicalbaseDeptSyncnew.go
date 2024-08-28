package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalbaseDeptSyncnew 直连科室上传
// alibaba.alihealth.medicalbase.dept.syncnew
//
// 直连科室上传接口
func AlibabaAlihealthMedicalbaseDeptSyncnew(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalbaseDeptSyncnewAPIRequest, resp *alihealth2.AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
