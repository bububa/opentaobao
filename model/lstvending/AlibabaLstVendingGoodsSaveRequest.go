package lstvending

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自动售卖机商品回传 APIRequest
alibaba.lst.vending.goods.save

零售通自动售卖机商品数据回流。
*/
type AlibabaLstVendingGoodsSaveRequest struct {
    model.Params

    // 商品信息
    goodsDTOList   []VendingGoodsDto 

}

func NewAlibabaLstVendingGoodsSaveRequest() *AlibabaLstVendingGoodsSaveRequest{
    return &AlibabaLstVendingGoodsSaveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstVendingGoodsSaveRequest) GetApiMethodName() string {
    return "alibaba.lst.vending.goods.save"
}

func (r AlibabaLstVendingGoodsSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstVendingGoodsSaveRequest) SetGoodsDTOList(goodsDTOList []VendingGoodsDto) error {
    r.goodsDTOList = goodsDTOList
    r.Set("goods_d_t_o_list", goodsDTOList)
    return nil
}

func (r AlibabaLstVendingGoodsSaveRequest) GetGoodsDTOList() []VendingGoodsDto {
    return r.goodsDTOList
}

