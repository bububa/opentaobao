package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticsexpressdeliveryresourcecreate 新建/更新配资源
// taobao.logistics.express.delivery.resource.create
//
// 新建/更新配资源
func Taobaologisticsexpressdeliveryresourcecreate(clt *core.SDKClient, req *ascp.TaobaologisticsexpressdeliveryresourcecreateAPIRequest, session string) (*ascp.TaobaologisticsexpressdeliveryresourcecreateAPIResponse, error) {
	var resp ascp.TaobaologisticsexpressdeliveryresourcecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
