package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// Alibabatcwmsoutboundloadboxcodecreate 创建箱号
// alibaba.tcwms.outbound.load.boxcode.create
//
// 创建箱号
func Alibabatcwmsoutboundloadboxcodecreate(clt *core.SDKClient, req *perfect.AlibabatcwmsoutboundloadboxcodecreateAPIRequest, session string) (*perfect.AlibabatcwmsoutboundloadboxcodecreateAPIResponse, error) {
	var resp perfect.AlibabatcwmsoutboundloadboxcodecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
