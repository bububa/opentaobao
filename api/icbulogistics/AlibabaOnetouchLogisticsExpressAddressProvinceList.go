package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// Alibabaonetouchlogisticsexpressaddressprovincelist 四级地址库-省
// alibaba.onetouch.logistics.express.address.province.list
//
// 四级地址库-省
func Alibabaonetouchlogisticsexpressaddressprovincelist(clt *core.SDKClient, req *icbulogistics.AlibabaonetouchlogisticsexpressaddressprovincelistAPIRequest, session string) (*icbulogistics.AlibabaonetouchlogisticsexpressaddressprovincelistAPIResponse, error) {
	var resp icbulogistics.AlibabaonetouchlogisticsexpressaddressprovincelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
