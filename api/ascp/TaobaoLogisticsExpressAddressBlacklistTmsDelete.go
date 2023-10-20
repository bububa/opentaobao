package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticsexpressaddressblacklisttmsdelete 上门取退可揽范围黑名单删除接口
// taobao.logistics.express.address.blacklist.tms.delete
//
// 上门取退可揽范围黑名单删除接口
func Taobaologisticsexpressaddressblacklisttmsdelete(clt *core.SDKClient, req *ascp.TaobaologisticsexpressaddressblacklisttmsdeleteAPIRequest, session string) (*ascp.TaobaologisticsexpressaddressblacklisttmsdeleteAPIResponse, error) {
	var resp ascp.TaobaologisticsexpressaddressblacklisttmsdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
