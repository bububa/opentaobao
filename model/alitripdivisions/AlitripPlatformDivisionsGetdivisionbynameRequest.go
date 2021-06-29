package alitripdivisions

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据中文名称与行政区划级别查询行政区划数据 APIRequest
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

func NewAlitripPlatformDivisionsGetdivisionbynameRequest() *AlitripPlatformDivisionsGetdivisionbynameRequest{
    return &AlitripPlatformDivisionsGetdivisionbynameRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripPlatformDivisionsGetdivisionbynameRequest) GetApiMethodName() string {
    return "alitrip.platform.divisions.getdivisionbyname"
}

func (r AlitripPlatformDivisionsGetdivisionbynameRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripPlatformDivisionsGetdivisionbynameRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlitripPlatformDivisionsGetdivisionbynameRequest) GetName() string {
    return r.name
}

func (r *AlitripPlatformDivisionsGetdivisionbynameRequest) SetLevel(level int64) error {
    r.level = level
    r.Set("level", level)
    return nil
}

func (r AlitripPlatformDivisionsGetdivisionbynameRequest) GetLevel() int64 {
    return r.level
}

