package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店城市数据获取接口 API请求
taobao.xhotel.city.get

引流API，对外提供酒店城市数据
*/
type TaobaoXhotelCityGetRequest struct {
    model.Params
    // 分页读取的开始下标,从0开始
    start   int64
    // 分页读取的城市个数，最小值为1，最大值为200
    count   int64
}

// 初始化TaobaoXhotelCityGetRequest对象
func NewTaobaoXhotelCityGetRequest() *TaobaoXhotelCityGetRequest{
    return &TaobaoXhotelCityGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelCityGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.city.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelCityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Start Setter
// 分页读取的开始下标,从0开始
func (r *TaobaoXhotelCityGetRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

// Start Getter
func (r TaobaoXhotelCityGetRequest) GetStart() int64 {
    return r.start
}
// Count Setter
// 分页读取的城市个数，最小值为1，最大值为200
func (r *TaobaoXhotelCityGetRequest) SetCount(count int64) error {
    r.count = count
    r.Set("count", count)
    return nil
}

// Count Getter
func (r TaobaoXhotelCityGetRequest) GetCount() int64 {
    return r.count
}
