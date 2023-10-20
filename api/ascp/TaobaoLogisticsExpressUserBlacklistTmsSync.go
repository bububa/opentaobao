package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticsexpressuserblacklisttmssync 上门取退用户黑名单同步
// taobao.logistics.express.user.blacklist.tms.sync
//
// 上门取退用户黑名单同步
func Taobaologisticsexpressuserblacklisttmssync(clt *core.SDKClient, req *ascp.TaobaologisticsexpressuserblacklisttmssyncAPIRequest, session string) (*ascp.TaobaologisticsexpressuserblacklisttmssyncAPIResponse, error) {
	var resp ascp.TaobaologisticsexpressuserblacklisttmssyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
