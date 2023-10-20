package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// Alibabatcwmsoutboundpickreceive 拣货接单
// alibaba.tcwms.outbound.pick.receive
//
// 拣货接单
func Alibabatcwmsoutboundpickreceive(clt *core.SDKClient, req *perfect.AlibabatcwmsoutboundpickreceiveAPIRequest, session string) (*perfect.AlibabatcwmsoutboundpickreceiveAPIResponse, error) {
	var resp perfect.AlibabatcwmsoutboundpickreceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
