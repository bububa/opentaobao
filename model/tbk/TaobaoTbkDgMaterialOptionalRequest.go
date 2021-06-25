package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-物料搜索 APIRequest
taobao.tbk.dg.material.optional

通用物料搜索API（导购）
*/
type TaobaoTbkDgMaterialOptionalRequest struct {
    model.Params

    // 商品筛选(特定媒体支持)-店铺dsr评分。筛选大于等于当前设置的店铺dsr评分的商品0-50000之间
    startDsr   int64 

    // 页大小，默认20，1~100
    pageSize   int64 

    // 第几页，默认：１
    pageNo   int64 

    // 链接形式：1：PC，2：无线，默认：１
    platform   int64 

    // 商品筛选-淘客佣金比率上限。如：1234表示12.34%
    endTkRate   int64 

    // 商品筛选-淘客佣金比率下限。如：1234表示12.34%
    startTkRate   int64 

    // 商品筛选-折扣价范围上限。单位：元
    endPrice   int64 

    // 商品筛选-折扣价范围下限。单位：元
    startPrice   int64 

    // 商品筛选-是否海外商品。true表示属于海外商品，false或不设置表示不限
    isOverseas   bool 

    // 商品筛选-是否天猫商品。true表示属于天猫商品，false或不设置表示不限
    isTmall   bool 

    // 排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price），匹配分（match）
    sort   string 

    // 商品筛选-所在地
    itemloc   string 

    // 商品筛选-后台类目ID。用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
    cat   string 

    // 商品筛选-查询词
    q   string 

    // 不传时默认物料id=2836；如果直接对消费者投放，可使用官方个性化算法优化的搜索物料id=17004
    materialId   int64 

    // 优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限
    hasCoupon   bool 

    // ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
    ip   string 

    // mm_xxx_xxx_12345678三段式的最后一段数字
    adzoneId   int64 

    // 商品筛选-是否包邮。true表示包邮，false或不设置表示不限
    needFreeShipment   bool 

    // 商品筛选-是否加入消费者保障。true表示加入，false或不设置表示不限
    needPrepay   bool 

    // 商品筛选(特定媒体支持)-成交转化是否高于行业均值。True表示大于等于，false或不设置表示不限
    includePayRate30   bool 

    // 商品筛选-好评率是否高于行业均值。True表示大于等于，false或不设置表示不限
    includeGoodRate   bool 

    // 商品筛选(特定媒体支持)-退款率是否低于行业均值。True表示大于等于，false或不设置表示不限
    includeRfdRate   bool 

    // 商品筛选-牛皮癣程度。取值：1不限，2无，3轻微
    npxLevel   int64 

    // 商品筛选-KA媒体淘客佣金比率上限。如：1234表示12.34%
    endKaTkRate   int64 

    // 商品筛选-KA媒体淘客佣金比率下限。如：1234表示12.34%
    startKaTkRate   int64 

    // 智能匹配-设备号加密类型：MD5
    deviceEncrypt   string 

    // 智能匹配-设备号加密后的值（MD5加密需32位小写）
    deviceValue   string 

    // 智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
    deviceType   string 

    // 锁佣结束时间
    lockRateEndTime   int64 

    // 锁佣开始时间
    lockRateStartTime   int64 

    // 本地化业务入参-LBS信息-经度
    longitude   string 

    // 本地化业务入参-LBS信息-纬度
    latitude   string 

    // 本地化业务入参-LBS信息-国标城市码，仅支持单个请求，请求饿了么卡券物料时，该字段必填。 （详细城市ID见：https://mo.m.taobao.com/page_2020010315120200508）
    cityCode   string 

    // 商家id，仅支持饿了么卡券商家ID，支持批量请求1-100以内，多个商家ID使用英文逗号分隔
    sellerIds   string 

    // 会员运营ID
    specialId   string 

    // 渠道关系ID，仅适用于渠道推广场景
    relationId   string 

    // 本地化业务入参-分页唯一标识，非首页的请求必传，值为上一页返回结果中的page_result_key字段值
    pageResultKey   string 

    // 人群ID，仅适用于物料评估场景material_id=41377
    ucrowdId   int64 

    // 物料评估-商品列表
    ucrowdRankItems   []Ucrowdrankitems 

    // 是否获取前N件佣金信息	0否，1是，其他值否
    getTopnRate   int64 

}

func NewTaobaoTbkDgMaterialOptionalRequest() *TaobaoTbkDgMaterialOptionalRequest{
    return &TaobaoTbkDgMaterialOptionalRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetApiMethodName() string {
    return "taobao.tbk.dg.material.optional"
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTbkDgMaterialOptionalRequest) SetStartDsr(startDsr int64) error {
    r.startDsr = startDsr
    r.Set("start_dsr", startDsr)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetStartDsr() int64 {
    return r.startDsr
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetPlatform(platform int64) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetPlatform() int64 {
    return r.platform
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetEndTkRate(endTkRate int64) error {
    r.endTkRate = endTkRate
    r.Set("end_tk_rate", endTkRate)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetEndTkRate() int64 {
    return r.endTkRate
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetStartTkRate(startTkRate int64) error {
    r.startTkRate = startTkRate
    r.Set("start_tk_rate", startTkRate)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetStartTkRate() int64 {
    return r.startTkRate
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetEndPrice(endPrice int64) error {
    r.endPrice = endPrice
    r.Set("end_price", endPrice)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetEndPrice() int64 {
    return r.endPrice
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetStartPrice(startPrice int64) error {
    r.startPrice = startPrice
    r.Set("start_price", startPrice)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetStartPrice() int64 {
    return r.startPrice
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetIsOverseas(isOverseas bool) error {
    r.isOverseas = isOverseas
    r.Set("is_overseas", isOverseas)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetIsOverseas() bool {
    return r.isOverseas
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetIsTmall(isTmall bool) error {
    r.isTmall = isTmall
    r.Set("is_tmall", isTmall)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetIsTmall() bool {
    return r.isTmall
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetSort(sort string) error {
    r.sort = sort
    r.Set("sort", sort)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetSort() string {
    return r.sort
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetItemloc(itemloc string) error {
    r.itemloc = itemloc
    r.Set("itemloc", itemloc)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetItemloc() string {
    return r.itemloc
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetCat(cat string) error {
    r.cat = cat
    r.Set("cat", cat)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetCat() string {
    return r.cat
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetQ(q string) error {
    r.q = q
    r.Set("q", q)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetQ() string {
    return r.q
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetMaterialId(materialId int64) error {
    r.materialId = materialId
    r.Set("material_id", materialId)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetMaterialId() int64 {
    return r.materialId
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetHasCoupon(hasCoupon bool) error {
    r.hasCoupon = hasCoupon
    r.Set("has_coupon", hasCoupon)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetHasCoupon() bool {
    return r.hasCoupon
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetIp() string {
    return r.ip
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetAdzoneId(adzoneId int64) error {
    r.adzoneId = adzoneId
    r.Set("adzone_id", adzoneId)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetAdzoneId() int64 {
    return r.adzoneId
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetNeedFreeShipment(needFreeShipment bool) error {
    r.needFreeShipment = needFreeShipment
    r.Set("need_free_shipment", needFreeShipment)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetNeedFreeShipment() bool {
    return r.needFreeShipment
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetNeedPrepay(needPrepay bool) error {
    r.needPrepay = needPrepay
    r.Set("need_prepay", needPrepay)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetNeedPrepay() bool {
    return r.needPrepay
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetIncludePayRate30(includePayRate30 bool) error {
    r.includePayRate30 = includePayRate30
    r.Set("include_pay_rate_30", includePayRate30)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetIncludePayRate30() bool {
    return r.includePayRate30
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetIncludeGoodRate(includeGoodRate bool) error {
    r.includeGoodRate = includeGoodRate
    r.Set("include_good_rate", includeGoodRate)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetIncludeGoodRate() bool {
    return r.includeGoodRate
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetIncludeRfdRate(includeRfdRate bool) error {
    r.includeRfdRate = includeRfdRate
    r.Set("include_rfd_rate", includeRfdRate)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetIncludeRfdRate() bool {
    return r.includeRfdRate
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetNpxLevel(npxLevel int64) error {
    r.npxLevel = npxLevel
    r.Set("npx_level", npxLevel)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetNpxLevel() int64 {
    return r.npxLevel
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetEndKaTkRate(endKaTkRate int64) error {
    r.endKaTkRate = endKaTkRate
    r.Set("end_ka_tk_rate", endKaTkRate)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetEndKaTkRate() int64 {
    return r.endKaTkRate
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetStartKaTkRate(startKaTkRate int64) error {
    r.startKaTkRate = startKaTkRate
    r.Set("start_ka_tk_rate", startKaTkRate)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetStartKaTkRate() int64 {
    return r.startKaTkRate
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetDeviceEncrypt(deviceEncrypt string) error {
    r.deviceEncrypt = deviceEncrypt
    r.Set("device_encrypt", deviceEncrypt)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetDeviceEncrypt() string {
    return r.deviceEncrypt
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetDeviceValue(deviceValue string) error {
    r.deviceValue = deviceValue
    r.Set("device_value", deviceValue)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetDeviceValue() string {
    return r.deviceValue
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetDeviceType() string {
    return r.deviceType
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetLockRateEndTime(lockRateEndTime int64) error {
    r.lockRateEndTime = lockRateEndTime
    r.Set("lock_rate_end_time", lockRateEndTime)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetLockRateEndTime() int64 {
    return r.lockRateEndTime
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetLockRateStartTime(lockRateStartTime int64) error {
    r.lockRateStartTime = lockRateStartTime
    r.Set("lock_rate_start_time", lockRateStartTime)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetLockRateStartTime() int64 {
    return r.lockRateStartTime
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetLongitude() string {
    return r.longitude
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetLatitude() string {
    return r.latitude
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetCityCode() string {
    return r.cityCode
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetSellerIds(sellerIds string) error {
    r.sellerIds = sellerIds
    r.Set("seller_ids", sellerIds)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetSellerIds() string {
    return r.sellerIds
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetSpecialId(specialId string) error {
    r.specialId = specialId
    r.Set("special_id", specialId)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetSpecialId() string {
    return r.specialId
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetRelationId(relationId string) error {
    r.relationId = relationId
    r.Set("relation_id", relationId)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetRelationId() string {
    return r.relationId
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetPageResultKey(pageResultKey string) error {
    r.pageResultKey = pageResultKey
    r.Set("page_result_key", pageResultKey)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetPageResultKey() string {
    return r.pageResultKey
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetUcrowdId(ucrowdId int64) error {
    r.ucrowdId = ucrowdId
    r.Set("ucrowd_id", ucrowdId)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetUcrowdId() int64 {
    return r.ucrowdId
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetUcrowdRankItems(ucrowdRankItems []Ucrowdrankitems) error {
    r.ucrowdRankItems = ucrowdRankItems
    r.Set("ucrowd_rank_items", ucrowdRankItems)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetUcrowdRankItems() []Ucrowdrankitems {
    return r.ucrowdRankItems
}

func (r *TaobaoTbkDgMaterialOptionalRequest) SetGetTopnRate(getTopnRate int64) error {
    r.getTopnRate = getTopnRate
    r.Set("get_topn_rate", getTopnRate)
    return nil
}

func (r TaobaoTbkDgMaterialOptionalRequest) GetGetTopnRate() int64 {
    return r.getTopnRate
}

