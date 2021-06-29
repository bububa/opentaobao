package alitripdivisions

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据中文名称与行政区划级别查询行政区划数据 API请求
alitrip.platform.divisions.getdivisionbyname

根据中文名称与行政区划级别查询行政区划数据
*/
type AlitripPlatformDivisionsGetdivisionbynameRequest struct {
    model.Params
    // 行政区划名称
    name   string
    // 行政区划级别ALL(0, "全部"),  	CONTINENT(1, "大洲"),  	COUNTRY(2, "国家"),  	PROVINCE(3, "省份"),  	CITY(4, "城市"),  	DISTRICT(5, "区县"),  	STREET(6, "街道")
    level   int64
}

// 初始化AlitripPlatformDivisionsGetdivisionbynameRequest对象
func NewAlitripPlatformDivisionsGetdivisionbynameRequest() *AlitripPlatformDivisionsGetdivisionbynameRequest{
    return &AlitripPlatformDivisionsGetdivisionbynameRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripPlatformDivisionsGetdivisionbynameRequest) GetApiMethodName() string {
    return "alitrip.platform.divisions.getdivisionbyname"
}

// IRequest interface 方法, 获取API参数
func (r AlitripPlatformDivisionsGetdivisionbynameRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 行政区划名称
func (r *AlitripPlatformDivisionsGetdivisionbynameRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlitripPlatformDivisionsGetdivisionbynameRequest) GetName() string {
    return r.name
}
// Level Setter
// 行政区划级别ALL(0, "全部"),  	CONTINENT(1, "大洲"),  	COUNTRY(2, "国家"),  	PROVINCE(3, "省份"),  	CITY(4, "城市"),  	DISTRICT(5, "区县"),  	STREET(6, "街道")
func (r *AlitripPlatformDivisionsGetdivisionbynameRequest) SetLevel(level int64) error {
    r.level = level
    r.Set("level", level)
    return nil
}

// Level Getter
func (r AlitripPlatformDivisionsGetdivisionbynameRequest) GetLevel() int64 {
    return r.level
}
