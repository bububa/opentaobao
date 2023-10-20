package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrVaequipmentList 获取企业冷链设备信息
// alibaba.alihealth.drug.kyt.dr.vaequipment.list
//
// 获取企业冷链设备信息
func AlibabaAlihealthDrugKytDrVaequipmentList(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrVaequipmentListAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrVaequipmentListAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytDrVaequipmentListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
