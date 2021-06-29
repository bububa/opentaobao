package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗预约门店列表查询 API请求
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

// 初始化AlibabaTaobaoMicdetailAlihealthQuerystoresRequest对象
func NewAlibabaTaobaoMicdetailAlihealthQuerystoresRequest() *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest{
    return &AlibabaTaobaoMicdetailAlihealthQuerystoresRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetApiMethodName() string {
    return "alibaba.taobao.micdetail.alihealth.querystores"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvType Setter
// 0不查询库存1查询库存
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetInvType(invType int64) error {
    r.invType = invType
    r.Set("inv_type", invType)
    return nil
}

// InvType Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetInvType() int64 {
    return r.invType
}
// PackageId Setter
// 包ID
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetPackageId(packageId int64) error {
    r.packageId = packageId
    r.Set("package_id", packageId)
    return nil
}

// PackageId Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetPackageId() int64 {
    return r.packageId
}
// PageSize Setter
// 页大小
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNum Setter
// 页码
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

// PageNum Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetPageNum() int64 {
    return r.pageNum
}
// Lon Setter
// 经度
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetLon(lon string) error {
    r.lon = lon
    r.Set("lon", lon)
    return nil
}

// Lon Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetLon() string {
    return r.lon
}
// Lat Setter
// 维度
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetLat(lat string) error {
    r.lat = lat
    r.Set("lat", lat)
    return nil
}

// Lat Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetLat() string {
    return r.lat
}
// ItemId Setter
// 商品ID
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetItemId() string {
    return r.itemId
}
// UserId Setter
// 用户ID
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetUserId() int64 {
    return r.userId
}
// DistrictId Setter
// 地区ID
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetDistrictId(districtId int64) error {
    r.districtId = districtId
    r.Set("district_id", districtId)
    return nil
}

// DistrictId Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetDistrictId() int64 {
    return r.districtId
}
// CityId Setter
// 城市ID
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetCityId(cityId int64) error {
    r.cityId = cityId
    r.Set("city_id", cityId)
    return nil
}

// CityId Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetCityId() int64 {
    return r.cityId
}
// ProvinceId Setter
// 省份ID
func (r *AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) SetProvinceId(provinceId int64) error {
    r.provinceId = provinceId
    r.Set("province_id", provinceId)
    return nil
}

// ProvinceId Getter
func (r AlibabaTaobaoMicdetailAlihealthQuerystoresRequest) GetProvinceId() int64 {
    return r.provinceId
}
