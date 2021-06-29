package axindata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
景点poi搜索-阿信 APIRequest
taobao.alitrip.travel.axin.poi.search

给阿信提供景点poi搜索
*/
type TaobaoAlitripTravelAxinPoiSearchRequest struct {
    model.Params

    // 搜索关键词
    keyWord   string 

}

func NewTaobaoAlitripTravelAxinPoiSearchRequest() *TaobaoAlitripTravelAxinPoiSearchRequest{
    return &TaobaoAlitripTravelAxinPoiSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelAxinPoiSearchRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.axin.poi.search"
}

func (r TaobaoAlitripTravelAxinPoiSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelAxinPoiSearchRequest) SetKeyWord(keyWord string) error {
    r.keyWord = keyWord
    r.Set("key_word", keyWord)
    return nil
}

func (r TaobaoAlitripTravelAxinPoiSearchRequest) GetKeyWord() string {
    return r.keyWord
}

