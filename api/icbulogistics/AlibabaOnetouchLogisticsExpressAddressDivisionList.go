package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// Alibabaonetouchlogisticsexpressaddressdivisionlist 四级地址库-区域
// alibaba.onetouch.logistics.express.address.division.list
//
// 四级地址库-区
func Alibabaonetouchlogisticsexpressaddressdivisionlist(clt *core.SDKClient, req *icbulogistics.AlibabaonetouchlogisticsexpressaddressdivisionlistAPIRequest, session string) (*icbulogistics.AlibabaonetouchlogisticsexpressaddressdivisionlistAPIResponse, error) {
	var resp icbulogistics.AlibabaonetouchlogisticsexpressaddressdivisionlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
