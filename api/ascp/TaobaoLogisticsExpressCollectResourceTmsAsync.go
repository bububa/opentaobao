package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticsexpresscollectresourcetmsasync 配服务商揽收能力同步接口
// taobao.logistics.express.collect.resource.tms.async
//
// 配服务商揽收能力同步接口
func Taobaologisticsexpresscollectresourcetmsasync(clt *core.SDKClient, req *ascp.TaobaologisticsexpresscollectresourcetmsasyncAPIRequest, session string) (*ascp.TaobaologisticsexpresscollectresourcetmsasyncAPIResponse, error) {
	var resp ascp.TaobaologisticsexpresscollectresourcetmsasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
