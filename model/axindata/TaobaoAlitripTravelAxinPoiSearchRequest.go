package axindata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
景点poi搜索-阿信 API请求
taobao.alitrip.travel.axin.poi.search

给阿信提供景点poi搜索
*/
type TaobaoAlitripTravelAxinPoiSearchRequest struct {
    model.Params
    // 搜索关键词
    _keyWord   string
}

// 初始化TaobaoAlitripTravelAxinPoiSearchRequest对象
func NewTaobaoAlitripTravelAxinPoiSearchRequest() *TaobaoAlitripTravelAxinPoiSearchRequest{
    return &TaobaoAlitripTravelAxinPoiSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinPoiSearchRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.axin.poi.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinPoiSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeyWord Setter
// 搜索关键词
func (r *TaobaoAlitripTravelAxinPoiSearchRequest) SetKeyWord(_keyWord string) error {
    r._keyWord = _keyWord
    r.Set("key_word", _keyWord)
    return nil
}

// KeyWord Getter
func (r TaobaoAlitripTravelAxinPoiSearchRequest) GetKeyWord() string {
    return r._keyWord
}
