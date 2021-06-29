package tuanhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店团购关联POI API请求
alitrip.tuan.hotel.relate.poi.get

查询酒店团购关联POI
*/
type AlitripTuanHotelRelatePoiGetRequest struct {
    model.Params
    // 关键字
    keywords   string
    // 行政区ID
    divisionId   int64
    // 类目ID：国内酒店套餐-201189402；国际酒店套餐-201188002；酒店餐饮美食-201214101；酒店服务-201214201；酒店客房优惠券-201214301
    catId   int64
}

// 初始化AlitripTuanHotelRelatePoiGetRequest对象
func NewAlitripTuanHotelRelatePoiGetRequest() *AlitripTuanHotelRelatePoiGetRequest{
    return &AlitripTuanHotelRelatePoiGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTuanHotelRelatePoiGetRequest) GetApiMethodName() string {
    return "alitrip.tuan.hotel.relate.poi.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTuanHotelRelatePoiGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Keywords Setter
// 关键字
func (r *AlitripTuanHotelRelatePoiGetRequest) SetKeywords(keywords string) error {
    r.keywords = keywords
    r.Set("keywords", keywords)
    return nil
}

// Keywords Getter
func (r AlitripTuanHotelRelatePoiGetRequest) GetKeywords() string {
    return r.keywords
}
// DivisionId Setter
// 行政区ID
func (r *AlitripTuanHotelRelatePoiGetRequest) SetDivisionId(divisionId int64) error {
    r.divisionId = divisionId
    r.Set("division_id", divisionId)
    return nil
}

// DivisionId Getter
func (r AlitripTuanHotelRelatePoiGetRequest) GetDivisionId() int64 {
    return r.divisionId
}
// CatId Setter
// 类目ID：国内酒店套餐-201189402；国际酒店套餐-201188002；酒店餐饮美食-201214101；酒店服务-201214201；酒店客房优惠券-201214301
func (r *AlitripTuanHotelRelatePoiGetRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

// CatId Getter
func (r AlitripTuanHotelRelatePoiGetRequest) GetCatId() int64 {
    return r.catId
}
