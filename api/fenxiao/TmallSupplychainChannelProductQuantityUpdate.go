package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Tmallsupplychainchannelproductquantityupdate 渠道无仓库存更新接口
// tmall.supplychain.channel.product.quantity.update
//
// 渠道无仓库存更新接口
func Tmallsupplychainchannelproductquantityupdate(clt *core.SDKClient, req *fenxiao.TmallsupplychainchannelproductquantityupdateAPIRequest, session string) (*fenxiao.TmallsupplychainchannelproductquantityupdateAPIResponse, error) {
	var resp fenxiao.TmallsupplychainchannelproductquantityupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
