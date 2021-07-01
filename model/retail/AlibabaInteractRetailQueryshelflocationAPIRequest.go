package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询货架和位置数据 API请求
alibaba.interact.retail.queryshelflocation

查询货架和位置数据
*/
type AlibabaInteractRetailQueryshelflocationAPIRequest struct {
    model.Params
    // 门店code
    _param0   string
}

// 初始化AlibabaInteractRetailQueryshelflocationAPIRequest对象
func NewAlibabaInteractRetailQueryshelflocationRequest() *AlibabaInteractRetailQueryshelflocationAPIRequest{
    return &AlibabaInteractRetailQueryshelflocationAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractRetailQueryshelflocationAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.retail.queryshelflocation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractRetailQueryshelflocationAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 门店code
func (r *AlibabaInteractRetailQueryshelflocationAPIRequest) SetParam0(_param0 string) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaInteractRetailQueryshelflocationAPIRequest) GetParam0() string {
    return r._param0
}
