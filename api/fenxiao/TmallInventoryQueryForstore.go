package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Tmallinventoryqueryforstore 查询后端商品仓库库存
// tmall.inventory.query.forstore
//
// 商家查询后端商品仓库库存
func Tmallinventoryqueryforstore(clt *core.SDKClient, req *fenxiao.TmallinventoryqueryforstoreAPIRequest, session string) (*fenxiao.TmallinventoryqueryforstoreAPIResponse, error) {
	var resp fenxiao.TmallinventoryqueryforstoreAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
