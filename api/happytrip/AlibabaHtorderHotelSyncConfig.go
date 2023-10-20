package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHtorderHotelSyncConfig 同步配置信息
// alibaba.htorder.hotel.sync.config
//
// 同步配置信息
func AlibabaHtorderHotelSyncConfig(clt *core.SDKClient, req *happytrip.AlibabaHtorderHotelSyncConfigAPIRequest, resp *happytrip.AlibabaHtorderHotelSyncConfigAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
