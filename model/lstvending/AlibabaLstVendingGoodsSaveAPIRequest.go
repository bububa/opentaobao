package lstvending

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自动售卖机商品回传 API请求
alibaba.lst.vending.goods.save

零售通自动售卖机商品数据回流。
*/
type AlibabaLstVendingGoodsSaveAPIRequest struct {
    model.Params
    // 商品信息
    _goodsDTOList   []VendingGoodsDto
}

// 初始化AlibabaLstVendingGoodsSaveAPIRequest对象
func NewAlibabaLstVendingGoodsSaveRequest() *AlibabaLstVendingGoodsSaveAPIRequest{
    return &AlibabaLstVendingGoodsSaveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingGoodsSaveAPIRequest) GetApiMethodName() string {
    return "alibaba.lst.vending.goods.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingGoodsSaveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GoodsDTOList Setter
// 商品信息
func (r *AlibabaLstVendingGoodsSaveAPIRequest) SetGoodsDTOList(_goodsDTOList []VendingGoodsDto) error {
    r._goodsDTOList = _goodsDTOList
    r.Set("goods_d_t_o_list", _goodsDTOList)
    return nil
}

// GoodsDTOList Getter
func (r AlibabaLstVendingGoodsSaveAPIRequest) GetGoodsDTOList() []VendingGoodsDto {
    return r._goodsDTOList
}
