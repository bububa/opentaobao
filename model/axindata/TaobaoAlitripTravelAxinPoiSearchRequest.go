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
type TaobaoAlitripTravelAxinPoiSearchAPIRequest struct {
    model.Params
    // 搜索关键词
    _keyWord   string
}

// 初始化TaobaoAlitripTravelAxinPoiSearchAPIRequest对象
func NewTaobaoAlitripTravelAxinPoiSearchRequest() *TaobaoAlitripTravelAxinPoiSearchAPIRequest{
    return &TaobaoAlitripTravelAxinPoiSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinPoiSearchAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.axin.poi.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinPoiSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeyWord Setter
// 搜索关键词
func (r *TaobaoAlitripTravelAxinPoiSearchAPIRequest) SetKeyWord(_keyWord string) error {
    r._keyWord = _keyWord
    r.Set("key_word", _keyWord)
    return nil
}

// KeyWord Getter
func (r TaobaoAlitripTravelAxinPoiSearchAPIRequest) GetKeyWord() string {
    return r._keyWord
}
