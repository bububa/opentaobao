package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品批量上下架接口 API请求
alibaba.icbu.product.batch.update.display

给国际站的三方服务商提供批量上下架接口
*/
type AlibabaIcbuProductBatchUpdateDisplayRequest struct {
    model.Params
    // on表示上架，off表示下架
    _newDisplay   string
    // 用逗号分隔的混淆id字符串
    _productIdList   string
}

// 初始化AlibabaIcbuProductBatchUpdateDisplayRequest对象
func NewAlibabaIcbuProductBatchUpdateDisplayRequest() *AlibabaIcbuProductBatchUpdateDisplayRequest{
    return &AlibabaIcbuProductBatchUpdateDisplayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductBatchUpdateDisplayRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.batch.update.display"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductBatchUpdateDisplayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NewDisplay Setter
// on表示上架，off表示下架
func (r *AlibabaIcbuProductBatchUpdateDisplayRequest) SetNewDisplay(_newDisplay string) error {
    r._newDisplay = _newDisplay
    r.Set("new_display", _newDisplay)
    return nil
}

// NewDisplay Getter
func (r AlibabaIcbuProductBatchUpdateDisplayRequest) GetNewDisplay() string {
    return r._newDisplay
}
// ProductIdList Setter
// 用逗号分隔的混淆id字符串
func (r *AlibabaIcbuProductBatchUpdateDisplayRequest) SetProductIdList(_productIdList string) error {
    r._productIdList = _productIdList
    r.Set("product_id_list", _productIdList)
    return nil
}

// ProductIdList Getter
func (r AlibabaIcbuProductBatchUpdateDisplayRequest) GetProductIdList() string {
    return r._productIdList
}
