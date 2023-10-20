package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpuopsupplierconsignordership 履约单商家仓发货结果回传服务
// alibaba.ascp.uop.supplier.consignorder.ship
//
// ERP通过该接口通知商家仓声明销售订单出库信息,支持履约单纬度全部发货的回传（目前不支持分批回传)
func Alibabaascpuopsupplierconsignordership(clt *core.SDKClient, req *ascpchannel.AlibabaascpuopsupplierconsignordershipAPIRequest, session string) (*ascpchannel.AlibabaascpuopsupplierconsignordershipAPIResponse, error) {
	var resp ascpchannel.AlibabaascpuopsupplierconsignordershipAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
