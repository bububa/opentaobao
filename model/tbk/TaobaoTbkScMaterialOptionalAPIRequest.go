package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScMaterialOptionalAPIRequest 淘宝客-服务商-物料搜索 API请求
// taobao.tbk.sc.material.optional
//
// 服务商使用。支持入参推广者对应的“推广位”、关键词和相关筛选条件，获取对应的物料信息和推广者对应的推广链接。
type TaobaoTbkScMaterialOptionalAPIRequest struct {
	model.Params
	// 物料评估-商品列表
	_ucrowdRankItems []Ucrowdrankitems
	// 排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price），匹配分（match）
	_sort string
	// 商品筛选-所在地
	_itemloc string
	// 商品筛选-后台类目ID。用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
	_cat string
	// 查询的关键词
	_q string
	// ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
	_ip string
	// 智能匹配-设备号加密后的值（MD5加密需32位小写）
	_deviceValue string
	// 智能匹配-设备号加密类型：MD5
	_deviceEncrypt string
	// 智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
	_deviceType string
	// 商家筛选-商家id，仅支持饿了么卡券商家id，支持批量请求1-100以内，多个商家id使用英文逗号分隔
	_sellerIds string
	// 本地化入参-LBS信息-国标城市码，仅支持单个请求，请求饿了么卡券物料时，该字段必填。 （详细城市ID见：https://mo.m.taobao.com/page_2020010315120200508）
	_cityCode string
	// 本地化入参-LBS信息-纬度
	_latitude string
	// 本地化入参-LBS信息-经度
	_longitude string
	// 渠道关系ID，仅适用于渠道推广场景
	_relationId string
	// 会员运营ID，仅适用于会员运营场景
	_specialId string
	// 本地化业务入参-分页唯一标识，非首页的请求必传，值为上一页返回结果中的page_result_key字段值
	_pageResultKey string
	// 1-动态ID转链场景，2-消费者比价场景（不填默认为1）
	_bizSceneId string
	// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）
	_promotionType string
	// 商品筛选(特定媒体支持)-店铺dsr评分。筛选大于等于当前设置的店铺dsr评分的商品0-50000之间
	_startDsr int64
	// 页大小，默认20，1~100
	_pageSize int64
	// 第几页，默认：１
	_pageNo int64
	// 链接形式-1：PC，2：无线，默认为１
	_platform int64
	// 商品筛选-淘客佣金比率上限。如：1234表示12.34%
	_endTkRate int64
	// 商品筛选-淘客佣金比率下限。如：1234表示12.34%
	_startTkRate int64
	// 商品筛选-折扣价范围上限。单位：元
	_endPrice int64
	// 商品筛选-折扣价范围上限。单位：元
	_startPrice int64
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneId int64
	// mm_xxx_xxx_xxx的第2段数字
	_siteId int64
	// 不传时默认物料id=2836。如果直接对消费者投放，可使用官方个性化算法优化的搜索物料id=17004
	_materialId int64
	// 商品筛选-牛皮癣程度。取值：1不限，2无，3轻微
	_npxLevel int64
	// 商品筛选-KA媒体淘客佣金率上限。如：1234表示12.34%
	_endKaTkRate int64
	// 商品筛选-KA媒体淘客佣金率下限。如：1234表示12.34%
	_startKaTkRate int64
	// 商品筛选-锁佣结束时间
	_lockRateEndTime int64
	// 商品筛选-锁佣开始时间
	_lockRateStartTime int64
	// 人群ID，仅适用于物料评估场景material_id=41377
	_ucrowdId int64
	// 是否获取前N件佣金信息，0否，1是，其他值否
	_getTopnRate int64
	// 商品筛选-是否海外商品。true表示属于海外商品，false或不设置表示不限
	_isOverseas bool
	// 商品筛选-是否天猫商品。true表示属于天猫商品，false或不设置表示不限
	_isTmall bool
	// 优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限
	_hasCoupon bool
	// 商品筛选(特定媒体支持)-退款率是否低于行业均值。True表示大于等于，false或不设置表示不限
	_includeRfdRate bool
	// 商品筛选(特定媒体支持)-好评率是否高于行业均值。True表示大于等于，false或不设置表示不限
	_includeGoodRate bool
	// 商品筛选(特定媒体支持)-成交转化是否高于行业均值。True表示大于等于，false或不设置表示不限
	_includePayRate30 bool
	// 商品筛选-是否加入消费者保障。true表示加入，false或不设置表示不限
	_needPrepay bool
	// 商品筛选-是否包邮。true表示包邮，false或不设置表示不限
	_needFreeShipment bool
}

// NewTaobaoTbkScMaterialOptionalRequest 初始化TaobaoTbkScMaterialOptionalAPIRequest对象
func NewTaobaoTbkScMaterialOptionalRequest() *TaobaoTbkScMaterialOptionalAPIRequest {
	return &TaobaoTbkScMaterialOptionalAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.material.optional"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUcrowdRankItems is UcrowdRankItems Setter
// 物料评估-商品列表
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetUcrowdRankItems(_ucrowdRankItems []Ucrowdrankitems) error {
	r._ucrowdRankItems = _ucrowdRankItems
	r.Set("ucrowd_rank_items", _ucrowdRankItems)
	return nil
}

// GetUcrowdRankItems UcrowdRankItems Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetUcrowdRankItems() []Ucrowdrankitems {
	return r._ucrowdRankItems
}

// SetSort is Sort Setter
// 排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price），匹配分（match）
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetSort() string {
	return r._sort
}

// SetItemloc is Itemloc Setter
// 商品筛选-所在地
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetItemloc(_itemloc string) error {
	r._itemloc = _itemloc
	r.Set("itemloc", _itemloc)
	return nil
}

// GetItemloc Itemloc Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetItemloc() string {
	return r._itemloc
}

// SetCat is Cat Setter
// 商品筛选-后台类目ID。用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetCat(_cat string) error {
	r._cat = _cat
	r.Set("cat", _cat)
	return nil
}

// GetCat Cat Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetCat() string {
	return r._cat
}

// SetQ is Q Setter
// 查询的关键词
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetQ(_q string) error {
	r._q = _q
	r.Set("q", _q)
	return nil
}

// GetQ Q Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetQ() string {
	return r._q
}

// SetIp is Ip Setter
// ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetIp() string {
	return r._ip
}

// SetDeviceValue is DeviceValue Setter
// 智能匹配-设备号加密后的值（MD5加密需32位小写）
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetDeviceValue(_deviceValue string) error {
	r._deviceValue = _deviceValue
	r.Set("device_value", _deviceValue)
	return nil
}

// GetDeviceValue DeviceValue Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetDeviceValue() string {
	return r._deviceValue
}

// SetDeviceEncrypt is DeviceEncrypt Setter
// 智能匹配-设备号加密类型：MD5
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetDeviceEncrypt(_deviceEncrypt string) error {
	r._deviceEncrypt = _deviceEncrypt
	r.Set("device_encrypt", _deviceEncrypt)
	return nil
}

// GetDeviceEncrypt DeviceEncrypt Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetDeviceEncrypt() string {
	return r._deviceEncrypt
}

// SetDeviceType is DeviceType Setter
// 智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetSellerIds is SellerIds Setter
// 商家筛选-商家id，仅支持饿了么卡券商家id，支持批量请求1-100以内，多个商家id使用英文逗号分隔
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetSellerIds(_sellerIds string) error {
	r._sellerIds = _sellerIds
	r.Set("seller_ids", _sellerIds)
	return nil
}

// GetSellerIds SellerIds Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetSellerIds() string {
	return r._sellerIds
}

// SetCityCode is CityCode Setter
// 本地化入参-LBS信息-国标城市码，仅支持单个请求，请求饿了么卡券物料时，该字段必填。 （详细城市ID见：https://mo.m.taobao.com/page_2020010315120200508）
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetLatitude is Latitude Setter
// 本地化入参-LBS信息-纬度
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetLongitude is Longitude Setter
// 本地化入参-LBS信息-经度
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetRelationId is RelationId Setter
// 渠道关系ID，仅适用于渠道推广场景
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetSpecialId is SpecialId Setter
// 会员运营ID，仅适用于会员运营场景
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// GetSpecialId SpecialId Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetSpecialId() string {
	return r._specialId
}

// SetPageResultKey is PageResultKey Setter
// 本地化业务入参-分页唯一标识，非首页的请求必传，值为上一页返回结果中的page_result_key字段值
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetPageResultKey(_pageResultKey string) error {
	r._pageResultKey = _pageResultKey
	r.Set("page_result_key", _pageResultKey)
	return nil
}

// GetPageResultKey PageResultKey Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetPageResultKey() string {
	return r._pageResultKey
}

// SetBizSceneId is BizSceneId Setter
// 1-动态ID转链场景，2-消费者比价场景（不填默认为1）
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetBizSceneId(_bizSceneId string) error {
	r._bizSceneId = _bizSceneId
	r.Set("biz_scene_id", _bizSceneId)
	return nil
}

// GetBizSceneId BizSceneId Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetBizSceneId() string {
	return r._bizSceneId
}

// SetPromotionType is PromotionType Setter
// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetPromotionType(_promotionType string) error {
	r._promotionType = _promotionType
	r.Set("promotion_type", _promotionType)
	return nil
}

// GetPromotionType PromotionType Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetPromotionType() string {
	return r._promotionType
}

// SetStartDsr is StartDsr Setter
// 商品筛选(特定媒体支持)-店铺dsr评分。筛选大于等于当前设置的店铺dsr评分的商品0-50000之间
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetStartDsr(_startDsr int64) error {
	r._startDsr = _startDsr
	r.Set("start_dsr", _startDsr)
	return nil
}

// GetStartDsr StartDsr Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetStartDsr() int64 {
	return r._startDsr
}

// SetPageSize is PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 第几页，默认：１
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPlatform is Platform Setter
// 链接形式-1：PC，2：无线，默认为１
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetEndTkRate is EndTkRate Setter
// 商品筛选-淘客佣金比率上限。如：1234表示12.34%
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetEndTkRate(_endTkRate int64) error {
	r._endTkRate = _endTkRate
	r.Set("end_tk_rate", _endTkRate)
	return nil
}

// GetEndTkRate EndTkRate Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetEndTkRate() int64 {
	return r._endTkRate
}

// SetStartTkRate is StartTkRate Setter
// 商品筛选-淘客佣金比率下限。如：1234表示12.34%
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetStartTkRate(_startTkRate int64) error {
	r._startTkRate = _startTkRate
	r.Set("start_tk_rate", _startTkRate)
	return nil
}

// GetStartTkRate StartTkRate Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetStartTkRate() int64 {
	return r._startTkRate
}

// SetEndPrice is EndPrice Setter
// 商品筛选-折扣价范围上限。单位：元
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetEndPrice(_endPrice int64) error {
	r._endPrice = _endPrice
	r.Set("end_price", _endPrice)
	return nil
}

// GetEndPrice EndPrice Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetEndPrice() int64 {
	return r._endPrice
}

// SetStartPrice is StartPrice Setter
// 商品筛选-折扣价范围上限。单位：元
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetStartPrice(_startPrice int64) error {
	r._startPrice = _startPrice
	r.Set("start_price", _startPrice)
	return nil
}

// GetStartPrice StartPrice Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetStartPrice() int64 {
	return r._startPrice
}

// SetAdzoneId is AdzoneId Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetSiteId is SiteId Setter
// mm_xxx_xxx_xxx的第2段数字
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetSiteId(_siteId int64) error {
	r._siteId = _siteId
	r.Set("site_id", _siteId)
	return nil
}

// GetSiteId SiteId Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetSiteId() int64 {
	return r._siteId
}

// SetMaterialId is MaterialId Setter
// 不传时默认物料id=2836。如果直接对消费者投放，可使用官方个性化算法优化的搜索物料id=17004
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetMaterialId(_materialId int64) error {
	r._materialId = _materialId
	r.Set("material_id", _materialId)
	return nil
}

// GetMaterialId MaterialId Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetMaterialId() int64 {
	return r._materialId
}

// SetNpxLevel is NpxLevel Setter
// 商品筛选-牛皮癣程度。取值：1不限，2无，3轻微
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetNpxLevel(_npxLevel int64) error {
	r._npxLevel = _npxLevel
	r.Set("npx_level", _npxLevel)
	return nil
}

// GetNpxLevel NpxLevel Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetNpxLevel() int64 {
	return r._npxLevel
}

// SetEndKaTkRate is EndKaTkRate Setter
// 商品筛选-KA媒体淘客佣金率上限。如：1234表示12.34%
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetEndKaTkRate(_endKaTkRate int64) error {
	r._endKaTkRate = _endKaTkRate
	r.Set("end_ka_tk_rate", _endKaTkRate)
	return nil
}

// GetEndKaTkRate EndKaTkRate Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetEndKaTkRate() int64 {
	return r._endKaTkRate
}

// SetStartKaTkRate is StartKaTkRate Setter
// 商品筛选-KA媒体淘客佣金率下限。如：1234表示12.34%
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetStartKaTkRate(_startKaTkRate int64) error {
	r._startKaTkRate = _startKaTkRate
	r.Set("start_ka_tk_rate", _startKaTkRate)
	return nil
}

// GetStartKaTkRate StartKaTkRate Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetStartKaTkRate() int64 {
	return r._startKaTkRate
}

// SetLockRateEndTime is LockRateEndTime Setter
// 商品筛选-锁佣结束时间
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetLockRateEndTime(_lockRateEndTime int64) error {
	r._lockRateEndTime = _lockRateEndTime
	r.Set("lock_rate_end_time", _lockRateEndTime)
	return nil
}

// GetLockRateEndTime LockRateEndTime Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetLockRateEndTime() int64 {
	return r._lockRateEndTime
}

// SetLockRateStartTime is LockRateStartTime Setter
// 商品筛选-锁佣开始时间
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetLockRateStartTime(_lockRateStartTime int64) error {
	r._lockRateStartTime = _lockRateStartTime
	r.Set("lock_rate_start_time", _lockRateStartTime)
	return nil
}

// GetLockRateStartTime LockRateStartTime Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetLockRateStartTime() int64 {
	return r._lockRateStartTime
}

// SetUcrowdId is UcrowdId Setter
// 人群ID，仅适用于物料评估场景material_id=41377
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetUcrowdId(_ucrowdId int64) error {
	r._ucrowdId = _ucrowdId
	r.Set("ucrowd_id", _ucrowdId)
	return nil
}

// GetUcrowdId UcrowdId Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetUcrowdId() int64 {
	return r._ucrowdId
}

// SetGetTopnRate is GetTopnRate Setter
// 是否获取前N件佣金信息，0否，1是，其他值否
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetGetTopnRate(_getTopnRate int64) error {
	r._getTopnRate = _getTopnRate
	r.Set("get_topn_rate", _getTopnRate)
	return nil
}

// GetGetTopnRate GetTopnRate Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetGetTopnRate() int64 {
	return r._getTopnRate
}

// SetIsOverseas is IsOverseas Setter
// 商品筛选-是否海外商品。true表示属于海外商品，false或不设置表示不限
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetIsOverseas(_isOverseas bool) error {
	r._isOverseas = _isOverseas
	r.Set("is_overseas", _isOverseas)
	return nil
}

// GetIsOverseas IsOverseas Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetIsOverseas() bool {
	return r._isOverseas
}

// SetIsTmall is IsTmall Setter
// 商品筛选-是否天猫商品。true表示属于天猫商品，false或不设置表示不限
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetIsTmall(_isTmall bool) error {
	r._isTmall = _isTmall
	r.Set("is_tmall", _isTmall)
	return nil
}

// GetIsTmall IsTmall Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetIsTmall() bool {
	return r._isTmall
}

// SetHasCoupon is HasCoupon Setter
// 优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetHasCoupon(_hasCoupon bool) error {
	r._hasCoupon = _hasCoupon
	r.Set("has_coupon", _hasCoupon)
	return nil
}

// GetHasCoupon HasCoupon Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetHasCoupon() bool {
	return r._hasCoupon
}

// SetIncludeRfdRate is IncludeRfdRate Setter
// 商品筛选(特定媒体支持)-退款率是否低于行业均值。True表示大于等于，false或不设置表示不限
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetIncludeRfdRate(_includeRfdRate bool) error {
	r._includeRfdRate = _includeRfdRate
	r.Set("include_rfd_rate", _includeRfdRate)
	return nil
}

// GetIncludeRfdRate IncludeRfdRate Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetIncludeRfdRate() bool {
	return r._includeRfdRate
}

// SetIncludeGoodRate is IncludeGoodRate Setter
// 商品筛选(特定媒体支持)-好评率是否高于行业均值。True表示大于等于，false或不设置表示不限
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetIncludeGoodRate(_includeGoodRate bool) error {
	r._includeGoodRate = _includeGoodRate
	r.Set("include_good_rate", _includeGoodRate)
	return nil
}

// GetIncludeGoodRate IncludeGoodRate Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetIncludeGoodRate() bool {
	return r._includeGoodRate
}

// SetIncludePayRate30 is IncludePayRate30 Setter
// 商品筛选(特定媒体支持)-成交转化是否高于行业均值。True表示大于等于，false或不设置表示不限
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetIncludePayRate30(_includePayRate30 bool) error {
	r._includePayRate30 = _includePayRate30
	r.Set("include_pay_rate_30", _includePayRate30)
	return nil
}

// GetIncludePayRate30 IncludePayRate30 Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetIncludePayRate30() bool {
	return r._includePayRate30
}

// SetNeedPrepay is NeedPrepay Setter
// 商品筛选-是否加入消费者保障。true表示加入，false或不设置表示不限
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetNeedPrepay(_needPrepay bool) error {
	r._needPrepay = _needPrepay
	r.Set("need_prepay", _needPrepay)
	return nil
}

// GetNeedPrepay NeedPrepay Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetNeedPrepay() bool {
	return r._needPrepay
}

// SetNeedFreeShipment is NeedFreeShipment Setter
// 商品筛选-是否包邮。true表示包邮，false或不设置表示不限
func (r *TaobaoTbkScMaterialOptionalAPIRequest) SetNeedFreeShipment(_needFreeShipment bool) error {
	r._needFreeShipment = _needFreeShipment
	r.Set("need_free_shipment", _needFreeShipment)
	return nil
}

// GetNeedFreeShipment NeedFreeShipment Getter
func (r TaobaoTbkScMaterialOptionalAPIRequest) GetNeedFreeShipment() bool {
	return r._needFreeShipment
}
