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
type AlibabaLstVendingGoodsSaveRequest struct {
    model.Params
    // 商品信息
    _goodsDTOList   []VendingGoodsDTO
}

// 初始化AlibabaLstVendingGoodsSaveRequest对象
func NewAlibabaLstVendingGoodsSaveRequest() *AlibabaLstVendingGoodsSaveRequest{
    return &AlibabaLstVendingGoodsSaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingGoodsSaveRequest) GetApiMethodName() string {
    return "alibaba.lst.vending.goods.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingGoodsSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GoodsDTOList Setter
// 商品信息
func (r *AlibabaLstVendingGoodsSaveRequest) SetGoodsDTOList(_goodsDTOList []VendingGoodsDTO) error {
    r._goodsDTOList = _goodsDTOList
    r.Set("goods_d_t_o_list", _goodsDTOList)
    return nil
}

// GoodsDTOList Getter
func (r AlibabaLstVendingGoodsSaveRequest) GetGoodsDTOList() []VendingGoodsDTO {
    return r._goodsDTOList
}
