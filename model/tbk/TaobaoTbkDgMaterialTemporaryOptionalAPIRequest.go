package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgMaterialTemporaryOptionalAPIRequest 淘宝客-推广者-物料搜索（临时接口） API请求
// taobao.tbk.dg.material.temporary.optional
//
// 通用物料搜索API（临时接口）
type TaobaoTbkDgMaterialTemporaryOptionalAPIRequest struct {
	model.Params
	// 物料评估-商品列表
	_ucrowdRankItems []Ucrowdrankitems
	// 排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price），匹配分（match）
	_sort string
	// 商品筛选-所在地
	_itemloc string
	// 商品筛选-后台类目ID。用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
	_cat string
	// 商品筛选-查询词
	_q string
	// ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
	_ip string
	// 智能匹配-设备号加密类型：MD5
	_deviceEncrypt string
	// 智能匹配-设备号加密后的值（MD5加密需32位小写）
	_deviceValue string
	// 智能匹配-设备类型：IMEI/IDFA/UTDID/OAID，（IOS推荐使用IDFA，android等推荐使用OAID）
	_deviceType string
	// 本地化业务入参-LBS信息-经度
	_longitude string
	// 本地化业务入参-LBS信息-纬度
	_latitude string
	// 本地化业务入参-LBS信息-国标城市码，仅支持单个请求，请求饿了么卡券物料时，该字段必填。 （详细城市ID见：https://mo.m.taobao.com/page_2020010315120200508）
	_cityCode string
	// 商家id，仅支持饿了么卡券商家ID，支持批量请求1-100以内，多个商家ID使用英文逗号分隔
	_sellerIds string
	// 会员运营ID
	_specialId string
	// 渠道关系ID，仅适用于渠道推广场景
	_relationId string
	// 本地化业务入参-分页唯一标识，非首页的请求必传，值为上一页返回结果中的page_result_key字段值
	_pageResultKey string
	// 商品筛选(特定媒体支持)-店铺dsr评分。筛选大于等于当前设置的店铺dsr评分的商品0-50000之间
	_startDsr int64
	// 页大小，默认20，1~100
	_pageSize int64
	// 第几页，默认：１
	_pageNo int64
	// 链接形式：1：PC，2：无线，默认：１
	_platform int64
	// 商品筛选-淘客佣金比率上限。如：1234表示12.34%
	_endTkRate int64
	// 商品筛选-淘客佣金比率下限。如：1234表示12.34%
	_startTkRate int64
	// 商品筛选-折扣价范围上限。单位：元
	_endPrice int64
	// 商品筛选-折扣价范围下限。单位：元
	_startPrice int64
	// 不传时默认物料id=2836；如果直接对消费者投放，可使用官方个性化算法优化的搜索物料id=17004
	_materialId int64
	// mm_xxx_xxx_12345678三段式的最后一段数字
	_adzoneId int64
	// 商品筛选-牛皮癣程度。取值：1不限，2无，3轻微
	_npxLevel int64
	// 商品筛选-KA媒体淘客佣金比率上限。如：1234表示12.34%
	_endKaTkRate int64
	// 商品筛选-KA媒体淘客佣金比率下限。如：1234表示12.34%
	_startKaTkRate int64
	// 锁佣结束时间
	_lockRateEndTime int64
	// 锁佣开始时间
	_lockRateStartTime int64
	// 人群ID，仅适用于物料评估场景material_id=41377
	_ucrowdId int64
	// 是否获取前N件佣金信息	0否，1是，其他值否
	_getTopnRate int64
	// 商品筛选-是否海外商品。true表示属于海外商品，false或不设置表示不限
	_isOverseas bool
	// 商品筛选-是否天猫商品。true表示属于天猫商品，false或不设置表示不限
	_isTmall bool
	// 优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限
	_hasCoupon bool
	// 商品筛选-是否包邮。true表示包邮，false或不设置表示不限
	_needFreeShipment bool
	// 商品筛选-是否加入消费者保障。true表示加入，false或不设置表示不限
	_needPrepay bool
	// 商品筛选(特定媒体支持)-成交转化是否高于行业均值。True表示大于等于，false或不设置表示不限
	_includePayRate30 bool
	// 商品筛选-好评率是否高于行业均值。True表示大于等于，false或不设置表示不限
	_includeGoodRate bool
	// 商品筛选(特定媒体支持)-退款率是否低于行业均值。True表示大于等于，false或不设置表示不限
	_includeRfdRate bool
}

// NewTaobaoTbkDgMaterialTemporaryOptionalRequest 初始化TaobaoTbkDgMaterialTemporaryOptionalAPIRequest对象
func NewTaobaoTbkDgMaterialTemporaryOptionalRequest() *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest {
	return &TaobaoTbkDgMaterialTemporaryOptionalAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.material.temporary.optional"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUcrowdRankItems is UcrowdRankItems Setter
// 物料评估-商品列表
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetUcrowdRankItems(_ucrowdRankItems []Ucrowdrankitems) error {
	r._ucrowdRankItems = _ucrowdRankItems
	r.Set("ucrowd_rank_items", _ucrowdRankItems)
	return nil
}

// GetUcrowdRankItems UcrowdRankItems Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetUcrowdRankItems() []Ucrowdrankitems {
	return r._ucrowdRankItems
}

// SetSort is Sort Setter
// 排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price），匹配分（match）
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetSort() string {
	return r._sort
}

// SetItemloc is Itemloc Setter
// 商品筛选-所在地
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetItemloc(_itemloc string) error {
	r._itemloc = _itemloc
	r.Set("itemloc", _itemloc)
	return nil
}

// GetItemloc Itemloc Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetItemloc() string {
	return r._itemloc
}

// SetCat is Cat Setter
// 商品筛选-后台类目ID。用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetCat(_cat string) error {
	r._cat = _cat
	r.Set("cat", _cat)
	return nil
}

// GetCat Cat Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetCat() string {
	return r._cat
}

// SetQ is Q Setter
// 商品筛选-查询词
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetQ(_q string) error {
	r._q = _q
	r.Set("q", _q)
	return nil
}

// GetQ Q Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetQ() string {
	return r._q
}

// SetIp is Ip Setter
// ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetIp() string {
	return r._ip
}

// SetDeviceEncrypt is DeviceEncrypt Setter
// 智能匹配-设备号加密类型：MD5
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetDeviceEncrypt(_deviceEncrypt string) error {
	r._deviceEncrypt = _deviceEncrypt
	r.Set("device_encrypt", _deviceEncrypt)
	return nil
}

// GetDeviceEncrypt DeviceEncrypt Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetDeviceEncrypt() string {
	return r._deviceEncrypt
}

// SetDeviceValue is DeviceValue Setter
// 智能匹配-设备号加密后的值（MD5加密需32位小写）
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetDeviceValue(_deviceValue string) error {
	r._deviceValue = _deviceValue
	r.Set("device_value", _deviceValue)
	return nil
}

// GetDeviceValue DeviceValue Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetDeviceValue() string {
	return r._deviceValue
}

// SetDeviceType is DeviceType Setter
// 智能匹配-设备类型：IMEI/IDFA/UTDID/OAID，（IOS推荐使用IDFA，android等推荐使用OAID）
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetLongitude is Longitude Setter
// 本地化业务入参-LBS信息-经度
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetLatitude is Latitude Setter
// 本地化业务入参-LBS信息-纬度
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetCityCode is CityCode Setter
// 本地化业务入参-LBS信息-国标城市码，仅支持单个请求，请求饿了么卡券物料时，该字段必填。 （详细城市ID见：https://mo.m.taobao.com/page_2020010315120200508）
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetSellerIds is SellerIds Setter
// 商家id，仅支持饿了么卡券商家ID，支持批量请求1-100以内，多个商家ID使用英文逗号分隔
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetSellerIds(_sellerIds string) error {
	r._sellerIds = _sellerIds
	r.Set("seller_ids", _sellerIds)
	return nil
}

// GetSellerIds SellerIds Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetSellerIds() string {
	return r._sellerIds
}

// SetSpecialId is SpecialId Setter
// 会员运营ID
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// GetSpecialId SpecialId Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetSpecialId() string {
	return r._specialId
}

// SetRelationId is RelationId Setter
// 渠道关系ID，仅适用于渠道推广场景
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetPageResultKey is PageResultKey Setter
// 本地化业务入参-分页唯一标识，非首页的请求必传，值为上一页返回结果中的page_result_key字段值
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetPageResultKey(_pageResultKey string) error {
	r._pageResultKey = _pageResultKey
	r.Set("page_result_key", _pageResultKey)
	return nil
}

// GetPageResultKey PageResultKey Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetPageResultKey() string {
	return r._pageResultKey
}

// SetStartDsr is StartDsr Setter
// 商品筛选(特定媒体支持)-店铺dsr评分。筛选大于等于当前设置的店铺dsr评分的商品0-50000之间
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetStartDsr(_startDsr int64) error {
	r._startDsr = _startDsr
	r.Set("start_dsr", _startDsr)
	return nil
}

// GetStartDsr StartDsr Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetStartDsr() int64 {
	return r._startDsr
}

// SetPageSize is PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 第几页，默认：１
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPlatform is Platform Setter
// 链接形式：1：PC，2：无线，默认：１
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetEndTkRate is EndTkRate Setter
// 商品筛选-淘客佣金比率上限。如：1234表示12.34%
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetEndTkRate(_endTkRate int64) error {
	r._endTkRate = _endTkRate
	r.Set("end_tk_rate", _endTkRate)
	return nil
}

// GetEndTkRate EndTkRate Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetEndTkRate() int64 {
	return r._endTkRate
}

// SetStartTkRate is StartTkRate Setter
// 商品筛选-淘客佣金比率下限。如：1234表示12.34%
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetStartTkRate(_startTkRate int64) error {
	r._startTkRate = _startTkRate
	r.Set("start_tk_rate", _startTkRate)
	return nil
}

// GetStartTkRate StartTkRate Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetStartTkRate() int64 {
	return r._startTkRate
}

// SetEndPrice is EndPrice Setter
// 商品筛选-折扣价范围上限。单位：元
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetEndPrice(_endPrice int64) error {
	r._endPrice = _endPrice
	r.Set("end_price", _endPrice)
	return nil
}

// GetEndPrice EndPrice Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetEndPrice() int64 {
	return r._endPrice
}

// SetStartPrice is StartPrice Setter
// 商品筛选-折扣价范围下限。单位：元
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetStartPrice(_startPrice int64) error {
	r._startPrice = _startPrice
	r.Set("start_price", _startPrice)
	return nil
}

// GetStartPrice StartPrice Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetStartPrice() int64 {
	return r._startPrice
}

// SetMaterialId is MaterialId Setter
// 不传时默认物料id=2836；如果直接对消费者投放，可使用官方个性化算法优化的搜索物料id=17004
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetMaterialId(_materialId int64) error {
	r._materialId = _materialId
	r.Set("material_id", _materialId)
	return nil
}

// GetMaterialId MaterialId Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetMaterialId() int64 {
	return r._materialId
}

// SetAdzoneId is AdzoneId Setter
// mm_xxx_xxx_12345678三段式的最后一段数字
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetNpxLevel is NpxLevel Setter
// 商品筛选-牛皮癣程度。取值：1不限，2无，3轻微
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetNpxLevel(_npxLevel int64) error {
	r._npxLevel = _npxLevel
	r.Set("npx_level", _npxLevel)
	return nil
}

// GetNpxLevel NpxLevel Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetNpxLevel() int64 {
	return r._npxLevel
}

// SetEndKaTkRate is EndKaTkRate Setter
// 商品筛选-KA媒体淘客佣金比率上限。如：1234表示12.34%
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetEndKaTkRate(_endKaTkRate int64) error {
	r._endKaTkRate = _endKaTkRate
	r.Set("end_ka_tk_rate", _endKaTkRate)
	return nil
}

// GetEndKaTkRate EndKaTkRate Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetEndKaTkRate() int64 {
	return r._endKaTkRate
}

// SetStartKaTkRate is StartKaTkRate Setter
// 商品筛选-KA媒体淘客佣金比率下限。如：1234表示12.34%
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetStartKaTkRate(_startKaTkRate int64) error {
	r._startKaTkRate = _startKaTkRate
	r.Set("start_ka_tk_rate", _startKaTkRate)
	return nil
}

// GetStartKaTkRate StartKaTkRate Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetStartKaTkRate() int64 {
	return r._startKaTkRate
}

// SetLockRateEndTime is LockRateEndTime Setter
// 锁佣结束时间
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetLockRateEndTime(_lockRateEndTime int64) error {
	r._lockRateEndTime = _lockRateEndTime
	r.Set("lock_rate_end_time", _lockRateEndTime)
	return nil
}

// GetLockRateEndTime LockRateEndTime Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetLockRateEndTime() int64 {
	return r._lockRateEndTime
}

// SetLockRateStartTime is LockRateStartTime Setter
// 锁佣开始时间
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetLockRateStartTime(_lockRateStartTime int64) error {
	r._lockRateStartTime = _lockRateStartTime
	r.Set("lock_rate_start_time", _lockRateStartTime)
	return nil
}

// GetLockRateStartTime LockRateStartTime Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetLockRateStartTime() int64 {
	return r._lockRateStartTime
}

// SetUcrowdId is UcrowdId Setter
// 人群ID，仅适用于物料评估场景material_id=41377
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetUcrowdId(_ucrowdId int64) error {
	r._ucrowdId = _ucrowdId
	r.Set("ucrowd_id", _ucrowdId)
	return nil
}

// GetUcrowdId UcrowdId Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetUcrowdId() int64 {
	return r._ucrowdId
}

// SetGetTopnRate is GetTopnRate Setter
// 是否获取前N件佣金信息	0否，1是，其他值否
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetGetTopnRate(_getTopnRate int64) error {
	r._getTopnRate = _getTopnRate
	r.Set("get_topn_rate", _getTopnRate)
	return nil
}

// GetGetTopnRate GetTopnRate Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetGetTopnRate() int64 {
	return r._getTopnRate
}

// SetIsOverseas is IsOverseas Setter
// 商品筛选-是否海外商品。true表示属于海外商品，false或不设置表示不限
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetIsOverseas(_isOverseas bool) error {
	r._isOverseas = _isOverseas
	r.Set("is_overseas", _isOverseas)
	return nil
}

// GetIsOverseas IsOverseas Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetIsOverseas() bool {
	return r._isOverseas
}

// SetIsTmall is IsTmall Setter
// 商品筛选-是否天猫商品。true表示属于天猫商品，false或不设置表示不限
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetIsTmall(_isTmall bool) error {
	r._isTmall = _isTmall
	r.Set("is_tmall", _isTmall)
	return nil
}

// GetIsTmall IsTmall Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetIsTmall() bool {
	return r._isTmall
}

// SetHasCoupon is HasCoupon Setter
// 优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetHasCoupon(_hasCoupon bool) error {
	r._hasCoupon = _hasCoupon
	r.Set("has_coupon", _hasCoupon)
	return nil
}

// GetHasCoupon HasCoupon Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetHasCoupon() bool {
	return r._hasCoupon
}

// SetNeedFreeShipment is NeedFreeShipment Setter
// 商品筛选-是否包邮。true表示包邮，false或不设置表示不限
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetNeedFreeShipment(_needFreeShipment bool) error {
	r._needFreeShipment = _needFreeShipment
	r.Set("need_free_shipment", _needFreeShipment)
	return nil
}

// GetNeedFreeShipment NeedFreeShipment Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetNeedFreeShipment() bool {
	return r._needFreeShipment
}

// SetNeedPrepay is NeedPrepay Setter
// 商品筛选-是否加入消费者保障。true表示加入，false或不设置表示不限
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetNeedPrepay(_needPrepay bool) error {
	r._needPrepay = _needPrepay
	r.Set("need_prepay", _needPrepay)
	return nil
}

// GetNeedPrepay NeedPrepay Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetNeedPrepay() bool {
	return r._needPrepay
}

// SetIncludePayRate30 is IncludePayRate30 Setter
// 商品筛选(特定媒体支持)-成交转化是否高于行业均值。True表示大于等于，false或不设置表示不限
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetIncludePayRate30(_includePayRate30 bool) error {
	r._includePayRate30 = _includePayRate30
	r.Set("include_pay_rate_30", _includePayRate30)
	return nil
}

// GetIncludePayRate30 IncludePayRate30 Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetIncludePayRate30() bool {
	return r._includePayRate30
}

// SetIncludeGoodRate is IncludeGoodRate Setter
// 商品筛选-好评率是否高于行业均值。True表示大于等于，false或不设置表示不限
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetIncludeGoodRate(_includeGoodRate bool) error {
	r._includeGoodRate = _includeGoodRate
	r.Set("include_good_rate", _includeGoodRate)
	return nil
}

// GetIncludeGoodRate IncludeGoodRate Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetIncludeGoodRate() bool {
	return r._includeGoodRate
}

// SetIncludeRfdRate is IncludeRfdRate Setter
// 商品筛选(特定媒体支持)-退款率是否低于行业均值。True表示大于等于，false或不设置表示不限
func (r *TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) SetIncludeRfdRate(_includeRfdRate bool) error {
	r._includeRfdRate = _includeRfdRate
	r.Set("include_rfd_rate", _includeRfdRate)
	return nil
}

// GetIncludeRfdRate IncludeRfdRate Getter
func (r TaobaoTbkDgMaterialTemporaryOptionalAPIRequest) GetIncludeRfdRate() bool {
	return r._includeRfdRate
}
