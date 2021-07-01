package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】基础信息获取接口：景点数据查询 API请求
taobao.alitrip.travel.baseinfo.scenics.get

商品发布辅助接口，用于飞猪度假或门票商品发布时 获取可用的景点（及景点下收费项目）信息列表。
*/
type TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest struct {
    model.Params
    // 城市名称
    _city   string
    // 景点简称
    _scenic   string
    // 景点id，可选。若传了该值，则查询结果中只会保留该id的景点，其余景点信息将被过滤
    _scenicId   int64
}

// 初始化TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest对象
func NewTaobaoAlitripTravelBaseinfoScenicsGetRequest() *TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest{
    return &TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.baseinfo.scenics.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// City Setter
// 城市名称
func (r *TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) GetCity() string {
    return r._city
}
// Scenic Setter
// 景点简称
func (r *TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) SetScenic(_scenic string) error {
    r._scenic = _scenic
    r.Set("scenic", _scenic)
    return nil
}

// Scenic Getter
func (r TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) GetScenic() string {
    return r._scenic
}
// ScenicId Setter
// 景点id，可选。若传了该值，则查询结果中只会保留该id的景点，其余景点信息将被过滤
func (r *TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) SetScenicId(_scenicId int64) error {
    r._scenicId = _scenicId
    r.Set("scenic_id", _scenicId)
    return nil
}

// ScenicId Getter
func (r TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest) GetScenicId() int64 {
    return r._scenicId
}
