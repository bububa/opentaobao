package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品批量上下架接口 APIRequest
alibaba.icbu.product.batch.update.display

给国际站的三方服务商提供批量上下架接口
*/
type AlibabaIcbuProductBatchUpdateDisplayRequest struct {
    model.Params

    // on表示上架，off表示下架
    newDisplay   string 

    // 用逗号分隔的混淆id字符串
    productIdList   string 

}

func NewAlibabaIcbuProductBatchUpdateDisplayRequest() *AlibabaIcbuProductBatchUpdateDisplayRequest{
    return &AlibabaIcbuProductBatchUpdateDisplayRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductBatchUpdateDisplayRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.batch.update.display"
}

func (r AlibabaIcbuProductBatchUpdateDisplayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductBatchUpdateDisplayRequest) SetNewDisplay(newDisplay string) error {
    r.newDisplay = newDisplay
    r.Set("new_display", newDisplay)
    return nil
}

func (r AlibabaIcbuProductBatchUpdateDisplayRequest) GetNewDisplay() string {
    return r.newDisplay
}

func (r *AlibabaIcbuProductBatchUpdateDisplayRequest) SetProductIdList(productIdList string) error {
    r.productIdList = productIdList
    r.Set("product_id_list", productIdList)
    return nil
}

func (r AlibabaIcbuProductBatchUpdateDisplayRequest) GetProductIdList() string {
    return r.productIdList
}

