package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗预约门店列表查询 APIRequest
alibaba.taobao.micdetail.alihealth.querystores

微信小程序疫苗预约门店列表查询
*/
type AlibabaTaobaoMicdetailAlihealthQuerystoresRequest struct {
    model.Params

    // 0不查询库存1查询库存
    invType   int64 

    // 包ID
    packageId   int64 

    // 页大小
    pageSize   int64 

    // 页码
    pageNum   int64 

    // 经度
    lon   string 

    // 维度
    lat   string 

    // 商品ID
    itemId   string 

    // 用户ID
    userId   int64 

    // 地区ID
    districtId   int64 

    // 城市ID
    cityId   int64 

    // 省份ID
    provinceId   int64 

}

func NewAlibabaTaobaoMicdetailAlihealthQuerystoresRequest() *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest{
    return &AlibabaTaobaoMicdetailAlihealthQuerystoresRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetApiMethodName() string {
    return "alibaba.taobao.micdetail.alihealth.querystores"
}

func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetInvType(invType int64) error {
    r.invType = invType
    r.Set("inv_type", invType)
    return nil
}

func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetInvType() int64 {
    return r.invType
}

func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetPackageId(packageId int64) error {
    r.packageId = packageId
    r.Set("package_id", packageId)
    return nil
}

func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetPackageId() int64 {
    return r.packageId
}

func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetPageNum() int64 {
    return r.pageNum
}

func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetLon(lon string) error {
    r.lon = lon
    r.Set("lon", lon)
    return nil
}

func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetLon() string {
    return r.lon
}

func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetLat(lat string) error {
    r.lat = lat
    r.Set("lat", lat)
    return nil
}

func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetLat() string {
    return r.lat
}

func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetItemId() string {
    return r.itemId
}

func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetUserId() int64 {
    return r.userId
}

func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetDistrictId(districtId int64) error {
    r.districtId = districtId
    r.Set("district_id", districtId)
    return nil
}

func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetDistrictId() int64 {
    return r.districtId
}

func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetCityId(cityId int64) error {
    r.cityId = cityId
    r.Set("city_id", cityId)
    return nil
}

func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetCityId() int64 {
    return r.cityId
}

func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetProvinceId(provinceId int64) error {
    r.provinceId = provinceId
    r.Set("province_id", provinceId)
    return nil
}

func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetProvinceId() int64 {
    return r.provinceId
}

