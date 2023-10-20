package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenwarehouseinfoquery 货主仓库资源查询接口
// taobao.qimen.warehouseinfo.query
//
// 货主仓库资源查询
func Taobaoqimenwarehouseinfoquery(clt *core.SDKClient, req *qimen.TaobaoqimenwarehouseinfoqueryAPIRequest, session string) (*qimen.TaobaoqimenwarehouseinfoqueryAPIResponse, error) {
	var resp qimen.TaobaoqimenwarehouseinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
