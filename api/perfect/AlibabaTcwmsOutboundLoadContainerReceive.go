package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// Alibabatcwmsoutboundloadcontainerreceive 装箱接单
// alibaba.tcwms.outbound.load.container.receive
//
// 装箱接单
func Alibabatcwmsoutboundloadcontainerreceive(clt *core.SDKClient, req *perfect.AlibabatcwmsoutboundloadcontainerreceiveAPIRequest, session string) (*perfect.AlibabatcwmsoutboundloadcontainerreceiveAPIResponse, error) {
	var resp perfect.AlibabatcwmsoutboundloadcontainerreceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
