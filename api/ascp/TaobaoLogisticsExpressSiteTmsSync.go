package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticsexpresssitetmssync 配服务商网点信息同步
// taobao.logistics.express.site.tms.sync
//
// 配服务商网点信息同步
func Taobaologisticsexpresssitetmssync(clt *core.SDKClient, req *ascp.TaobaologisticsexpresssitetmssyncAPIRequest, session string) (*ascp.TaobaologisticsexpresssitetmssyncAPIResponse, error) {
	var resp ascp.TaobaologisticsexpresssitetmssyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
