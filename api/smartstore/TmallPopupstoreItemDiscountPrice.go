package smartstore

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/smartstore"
)

/* 
商品优惠价格查询 
tmall.popupstore.item.discount.price

商品优惠价格查询
*/
func TmallPopupstoreItemDiscountPrice(clt *core.SDKClient, req *smartstore.TmallPopupstoreItemDiscountPriceRequest, session string) (*smartstore.TmallPopupstoreItemDiscountPriceAPIResponse, error) {
    var resp smartstore.TmallPopupstoreItemDiscountPriceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
