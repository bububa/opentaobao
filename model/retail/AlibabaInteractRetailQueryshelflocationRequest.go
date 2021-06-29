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
type AlibabaInteractRetailQueryshelflocationRequest struct {
    model.Params
    // 门店code
    _param0   string
}

// 初始化AlibabaInteractRetailQueryshelflocationRequest对象
func NewAlibabaInteractRetailQueryshelflocationRequest() *AlibabaInteractRetailQueryshelflocationRequest{
    return &AlibabaInteractRetailQueryshelflocationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractRetailQueryshelflocationRequest) GetApiMethodName() string {
    return "alibaba.interact.retail.queryshelflocation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractRetailQueryshelflocationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 门店code
func (r *AlibabaInteractRetailQueryshelflocationRequest) SetParam0(_param0 string) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaInteractRetailQueryshelflocationRequest) GetParam0() string {
    return r._param0
}
