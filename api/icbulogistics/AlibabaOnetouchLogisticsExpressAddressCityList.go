package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// Alibabaonetouchlogisticsexpressaddresscitylist 四级地址库-市
// alibaba.onetouch.logistics.express.address.city.list
//
// 四级地址库-市
func Alibabaonetouchlogisticsexpressaddresscitylist(clt *core.SDKClient, req *icbulogistics.AlibabaonetouchlogisticsexpressaddresscitylistAPIRequest, session string) (*icbulogistics.AlibabaonetouchlogisticsexpressaddresscitylistAPIResponse, error) {
	var resp icbulogistics.AlibabaonetouchlogisticsexpressaddresscitylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
