package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrvaequipmentlist 获取企业冷链设备信息
// alibaba.alihealth.drug.kyt.dr.vaequipment.list
//
// 获取企业冷链设备信息
func Alibabaalihealthdrugkytdrvaequipmentlist(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrvaequipmentlistAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrvaequipmentlistAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrvaequipmentlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
