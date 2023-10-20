package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticsexpresscapacitytmsasync 上门取退产能信息同步/更新
// taobao.logistics.express.capacity.tms.async
//
// 上门取退产能信息同步/更新
func Taobaologisticsexpresscapacitytmsasync(clt *core.SDKClient, req *ascp.TaobaologisticsexpresscapacitytmsasyncAPIRequest, session string) (*ascp.TaobaologisticsexpresscapacitytmsasyncAPIResponse, error) {
	var resp ascp.TaobaologisticsexpresscapacitytmsasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
