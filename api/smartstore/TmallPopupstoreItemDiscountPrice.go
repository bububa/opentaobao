package smartstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/smartstore"
)

// Tmallpopupstoreitemdiscountprice 商品优惠价格查询
// tmall.popupstore.item.discount.price
//
// 商品优惠价格查询
func Tmallpopupstoreitemdiscountprice(clt *core.SDKClient, req *smartstore.TmallpopupstoreitemdiscountpriceAPIRequest, session string) (*smartstore.TmallpopupstoreitemdiscountpriceAPIResponse, error) {
	var resp smartstore.TmallpopupstoreitemdiscountpriceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
