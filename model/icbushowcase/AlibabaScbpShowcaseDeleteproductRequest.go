package icbushowcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除橱窗商品 APIRequest
alibaba.scbp.showcase.deleteproduct

批量删除橱窗商品
*/
type AlibabaScbpShowcaseDeleteproductRequest struct {
    model.Params

    // 橱窗idList
    windowIdList   []int64 

}

func NewAlibabaScbpShowcaseDeleteproductRequest() *AlibabaScbpShowcaseDeleteproductRequest{
    return &AlibabaScbpShowcaseDeleteproductRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpShowcaseDeleteproductRequest) GetApiMethodName() string {
    return "alibaba.scbp.showcase.deleteproduct"
}

func (r AlibabaScbpShowcaseDeleteproductRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpShowcaseDeleteproductRequest) SetWindowIdList(windowIdList []int64) error {
    r.windowIdList = windowIdList
    r.Set("window_id_list", windowIdList)
    return nil
}

func (r AlibabaScbpShowcaseDeleteproductRequest) GetWindowIdList() []int64 {
    return r.windowIdList
}

