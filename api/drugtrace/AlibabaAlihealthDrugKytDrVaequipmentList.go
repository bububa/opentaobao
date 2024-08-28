package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrVaequipmentList 获取企业冷链设备信息
// alibaba.alihealth.drug.kyt.dr.vaequipment.list
//
// 获取企业冷链设备信息
func AlibabaAlihealthDrugKytDrVaequipmentList(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrVaequipmentListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
