package medicalbase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/medicalbase"
)

// AlibabaAlihealthMedicalbaseOrderStatusSync 号源直连订单状态同步接口
// alibaba.alihealth.medicalbase.order.status.sync
//
// 互联网医院isv批量通过接口批量导入
func AlibabaAlihealthMedicalbaseOrderStatusSync(clt *core.SDKClient, req *medicalbase.AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest, session string) (*medicalbase.AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse, error) {
	var resp medicalbase.AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
