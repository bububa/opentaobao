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
type TaobaoAlitripTravelBaseinfoScenicsGetRequest struct {
    model.Params
    // 城市名称
    city   string
    // 景点简称
    scenic   string
    // 景点id，可选。若传了该值，则查询结果中只会保留该id的景点，其余景点信息将被过滤
    scenicId   int64
}

// 初始化TaobaoAlitripTravelBaseinfoScenicsGetRequest对象
func NewTaobaoAlitripTravelBaseinfoScenicsGetRequest() *TaobaoAlitripTravelBaseinfoScenicsGetRequest{
    return &TaobaoAlitripTravelBaseinfoScenicsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelBaseinfoScenicsGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.baseinfo.scenics.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelBaseinfoScenicsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// City Setter
// 城市名称
func (r *TaobaoAlitripTravelBaseinfoScenicsGetRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r TaobaoAlitripTravelBaseinfoScenicsGetRequest) GetCity() string {
    return r.city
}
// Scenic Setter
// 景点简称
func (r *TaobaoAlitripTravelBaseinfoScenicsGetRequest) SetScenic(scenic string) error {
    r.scenic = scenic
    r.Set("scenic", scenic)
    return nil
}

// Scenic Getter
func (r TaobaoAlitripTravelBaseinfoScenicsGetRequest) GetScenic() string {
    return r.scenic
}
// ScenicId Setter
// 景点id，可选。若传了该值，则查询结果中只会保留该id的景点，其余景点信息将被过滤
func (r *TaobaoAlitripTravelBaseinfoScenicsGetRequest) SetScenicId(scenicId int64) error {
    r.scenicId = scenicId
    r.Set("scenic_id", scenicId)
    return nil
}

// ScenicId Getter
func (r TaobaoAlitripTravelBaseinfoScenicsGetRequest) GetScenicId() int64 {
    return r.scenicId
}
