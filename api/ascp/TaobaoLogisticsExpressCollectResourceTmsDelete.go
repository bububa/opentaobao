package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticsexpresscollectresourcetmsdelete 上门取退可揽范围删除
// taobao.logistics.express.collect.resource.tms.delete
//
// 上门取退可揽范围删除
func Taobaologisticsexpresscollectresourcetmsdelete(clt *core.SDKClient, req *ascp.TaobaologisticsexpresscollectresourcetmsdeleteAPIRequest, session string) (*ascp.TaobaologisticsexpresscollectresourcetmsdeleteAPIResponse, error) {
	var resp ascp.TaobaologisticsexpresscollectresourcetmsdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
