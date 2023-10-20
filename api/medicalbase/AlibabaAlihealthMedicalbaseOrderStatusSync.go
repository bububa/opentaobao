package medicalbase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/medicalbase"
)

// AlibabaAlihealthMedicalbaseOrderStatusSync 号源直连订单状态同步接口
// alibaba.alihealth.medicalbase.order.status.sync
//
// 互联网医院isv批量通过接口批量导入
func AlibabaAlihealthMedicalbaseOrderStatusSync(clt *core.SDKClient, req *medicalbase.AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest, resp *medicalbase.AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
