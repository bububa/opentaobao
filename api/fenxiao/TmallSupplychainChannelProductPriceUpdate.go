package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Tmallsupplychainchannelproductpriceupdate 渠道价格更新接口
// tmall.supplychain.channel.product.price.update
//
// 更新渠道产品价格
func Tmallsupplychainchannelproductpriceupdate(clt *core.SDKClient, req *fenxiao.TmallsupplychainchannelproductpriceupdateAPIRequest, session string) (*fenxiao.TmallsupplychainchannelproductpriceupdateAPIResponse, error) {
	var resp fenxiao.TmallsupplychainchannelproductpriceupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
