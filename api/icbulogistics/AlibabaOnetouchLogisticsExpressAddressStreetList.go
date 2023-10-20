package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// Alibabaonetouchlogisticsexpressaddressstreetlist 四级地址库-街道
// alibaba.onetouch.logistics.express.address.street.list
//
// 四级地址库-街道模糊查询
func Alibabaonetouchlogisticsexpressaddressstreetlist(clt *core.SDKClient, req *icbulogistics.AlibabaonetouchlogisticsexpressaddressstreetlistAPIRequest, session string) (*icbulogistics.AlibabaonetouchlogisticsexpressaddressstreetlistAPIResponse, error) {
	var resp icbulogistics.AlibabaonetouchlogisticsexpressaddressstreetlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
