package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticsexpressaddressblacklisttmsasync 上门取退可揽范围黑名单同步/更新
// taobao.logistics.express.address.blacklist.tms.async
//
// 上门取退可揽范围黑名单同步/更新
func Taobaologisticsexpressaddressblacklisttmsasync(clt *core.SDKClient, req *ascp.TaobaologisticsexpressaddressblacklisttmsasyncAPIRequest, session string) (*ascp.TaobaologisticsexpressaddressblacklisttmsasyncAPIResponse, error) {
	var resp ascp.TaobaologisticsexpressaddressblacklisttmsasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
