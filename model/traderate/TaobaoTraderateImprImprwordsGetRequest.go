package traderate

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
评价大家印象印象短语接口 APIRequest
taobao.traderate.impr.imprwords.get

根据淘宝后台类目的一级类目和叶子类目
*/
type TaobaoTraderateImprImprwordsGetRequest struct {
    model.Params

    // 淘宝一级类目id
    catRootId   int64 

    // 淘宝叶子类目id
    catLeafId   int64 

}

func NewTaobaoTraderateImprImprwordsGetRequest() *TaobaoTraderateImprImprwordsGetRequest{
    return &TaobaoTraderateImprImprwordsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTraderateImprImprwordsGetRequest) GetApiMethodName() string {
    return "taobao.traderate.impr.imprwords.get"
}

func (r TaobaoTraderateImprImprwordsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTraderateImprImprwordsGetRequest) SetCatRootId(catRootId int64) error {
    r.catRootId = catRootId
    r.Set("cat_root_id", catRootId)
    return nil
}

func (r TaobaoTraderateImprImprwordsGetRequest) GetCatRootId() int64 {
    return r.catRootId
}

func (r *TaobaoTraderateImprImprwordsGetRequest) SetCatLeafId(catLeafId int64) error {
    r.catLeafId = catLeafId
    r.Set("cat_leaf_id", catLeafId)
    return nil
}

func (r TaobaoTraderateImprImprwordsGetRequest) GetCatLeafId() int64 {
    return r.catLeafId
}

