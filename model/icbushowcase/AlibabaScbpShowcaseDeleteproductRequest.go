package icbushowcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除橱窗商品 API请求
alibaba.scbp.showcase.deleteproduct

批量删除橱窗商品
*/
type AlibabaScbpShowcaseDeleteproductRequest struct {
    model.Params
    // 橱窗idList
    _windowIdList   []int64
}

// 初始化AlibabaScbpShowcaseDeleteproductRequest对象
func NewAlibabaScbpShowcaseDeleteproductRequest() *AlibabaScbpShowcaseDeleteproductRequest{
    return &AlibabaScbpShowcaseDeleteproductRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseDeleteproductRequest) GetApiMethodName() string {
    return "alibaba.scbp.showcase.deleteproduct"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseDeleteproductRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WindowIdList Setter
// 橱窗idList
func (r *AlibabaScbpShowcaseDeleteproductRequest) SetWindowIdList(_windowIdList []int64) error {
    r._windowIdList = _windowIdList
    r.Set("window_id_list", _windowIdList)
    return nil
}

// WindowIdList Getter
func (r AlibabaScbpShowcaseDeleteproductRequest) GetWindowIdList() []int64 {
    return r._windowIdList
}
