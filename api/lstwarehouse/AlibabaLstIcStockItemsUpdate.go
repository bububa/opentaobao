package lstwarehouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstwarehouse"
)

// Alibabalsticstockitemsupdate 零售通经销商商品库存设置
// alibaba.lst.ic.stock.items.update
//
// 零售通经销商商品库存设置
func Alibabalsticstockitemsupdate(clt *core.SDKClient, req *lstwarehouse.AlibabalsticstockitemsupdateAPIRequest, session string) (*lstwarehouse.AlibabalsticstockitemsupdateAPIResponse, error) {
	var resp lstwarehouse.AlibabalsticstockitemsupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
