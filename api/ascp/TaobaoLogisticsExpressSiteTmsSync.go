package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressSiteTmsSync 配服务商网点信息同步
// taobao.logistics.express.site.tms.sync
//
// 配服务商网点信息同步
func TaobaoLogisticsExpressSiteTmsSync(clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressSiteTmsSyncAPIRequest, session string) (*ascp.TaobaoLogisticsExpressSiteTmsSyncAPIResponse, error) {
	var resp ascp.TaobaoLogisticsExpressSiteTmsSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
