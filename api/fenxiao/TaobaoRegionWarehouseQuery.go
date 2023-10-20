package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoregionwarehousequery 查询仓库覆盖范围
// taobao.region.warehouse.query
//
// 查询仓库覆盖范围
func Taobaoregionwarehousequery(clt *core.SDKClient, req *fenxiao.TaobaoregionwarehousequeryAPIRequest, session string) (*fenxiao.TaobaoregionwarehousequeryAPIResponse, error) {
	var resp fenxiao.TaobaoregionwarehousequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
