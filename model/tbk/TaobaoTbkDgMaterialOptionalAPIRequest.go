package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgmaterialoptionalAPIRequest 淘宝客-推广者-物料搜索 API请求
// taobao.tbk.dg.material.optional
//
// 通用物料搜索API（导购）
type TaobaotbkdgmaterialoptionalAPIRequest struct {
	model.Params
	// 物料评估-商品列表
	_ucrowdrankitems []Ucrowdrankitems
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
	_deviceencrypt string
	// 智能匹配-设备号加密后的值（MD5加密需32位小写）
	_devicevalue string
	// 智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
	_devicetype string
	// 本地化业务入参-LBS信息-经度
	_longitude string
	// 本地化业务入参-LBS信息-纬度
	_latitude string
	// 本地化业务入参-LBS信息-国标城市码，仅支持单个请求，请求饿了么卡券物料时，该字段必填。 （详细城市ID见：https://mo.m.taobao.com/page_2020010315120200508）
	_citycode string
	// 商家id，仅支持饿了么卡券商家ID，支持批量请求1-100以内，多个商家ID使用英文逗号分隔
	_sellerids string
	// 会员运营ID
	_specialid string
	// 渠道关系ID，仅适用于渠道推广场景
	_relationid string
	// 本地化业务入参-分页唯一标识，非首页的请求必传，值为上一页返回结果中的page_result_key字段值
	_pageresultkey string
	// 1-动态ID转链场景，2-消费者比价场景（不填默认为1）
	_bizsceneid string
	// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）
	_promotiontype string
	// 商品筛选(特定媒体支持)-店铺dsr评分。筛选大于等于当前设置的店铺dsr评分的商品0-50000之间
	_startdsr int64
	// 页大小，默认20，1~100
	_pagesize int64
	// 第几页，默认：１
	_pageno int64
	// 链接形式：1：PC，2：无线，默认：１
	_platform int64
	// 商品筛选-淘客佣金比率上限。如：1234表示12.34%
	_endtkrate int64
	// 商品筛选-淘客佣金比率下限。如：1234表示12.34%
	_starttkrate int64
	// 商品筛选-折扣价范围上限。单位：元
	_endprice int64
	// 商品筛选-折扣价范围下限。单位：元
	_startprice int64
	// 不传时默认物料id=2836；如果直接对消费者投放，可使用官方个性化算法优化的搜索物料id=17004
	_materialid int64
	// mm_xxx_xxx_12345678三段式的最后一段数字
	_adzoneid int64
	// 商品筛选-牛皮癣程度。取值：1不限，2无，3轻微
	_npxlevel int64
	// 商品筛选-KA媒体淘客佣金比率上限。如：1234表示12.34%
	_endkatkrate int64
	// 商品筛选-KA媒体淘客佣金比率下限。如：1234表示12.34%
	_startkatkrate int64
	// 锁佣结束时间
	_lockrateendtime int64
	// 锁佣开始时间
	_lockratestarttime int64
	// 人群ID，仅适用于物料评估场景material_id=41377
	_ucrowdid int64
	// 是否获取前N件佣金信息	0否，1是，其他值否
	_gettopnrate int64
	// 商品筛选-是否海外商品。true表示属于海外商品，false或不设置表示不限
	_isoverseas bool
	// 商品筛选-是否天猫商品。true表示属于天猫商品，false或不设置表示不限
	_istmall bool
	// 优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限
	_hascoupon bool
	// 商品筛选-是否包邮。true表示包邮，false或不设置表示不限
	_needfreeshipment bool
	// 商品筛选-是否加入消费者保障。true表示加入，false或不设置表示不限
	_needprepay bool
	// 商品筛选(特定媒体支持)-成交转化是否高于行业均值。True表示大于等于，false或不设置表示不限
	_includepayrate30 bool
	// 商品筛选-好评率是否高于行业均值。True表示大于等于，false或不设置表示不限
	_includegoodrate bool
	// 商品筛选(特定媒体支持)-退款率是否低于行业均值。True表示大于等于，false或不设置表示不限
	_includerfdrate bool
}

// NewTaobaotbkdgmaterialoptionalRequest 初始化TaobaotbkdgmaterialoptionalAPIRequest对象
func NewTaobaotbkdgmaterialoptionalRequest() *TaobaotbkdgmaterialoptionalAPIRequest {
	return &TaobaotbkdgmaterialoptionalAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.material.optional"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUcrowdrankitems is Ucrowdrankitems Setter
// 物料评估-商品列表
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetUcrowdrankitems(_ucrowdrankitems []Ucrowdrankitems) error {
	r._ucrowdrankitems = _ucrowdrankitems
	r.Set("ucrowd_rank_items", _ucrowdrankitems)
	return nil
}

// GetUcrowdrankitems Ucrowdrankitems Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetUcrowdrankitems() []Ucrowdrankitems {
	return r._ucrowdrankitems
}

// SetSort is Sort Setter
// 排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price），匹配分（match）
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetSort() string {
	return r._sort
}

// SetItemloc is Itemloc Setter
// 商品筛选-所在地
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetItemloc(_itemloc string) error {
	r._itemloc = _itemloc
	r.Set("itemloc", _itemloc)
	return nil
}

// GetItemloc Itemloc Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetItemloc() string {
	return r._itemloc
}

// SetCat is Cat Setter
// 商品筛选-后台类目ID。用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetCat(_cat string) error {
	r._cat = _cat
	r.Set("cat", _cat)
	return nil
}

// GetCat Cat Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetCat() string {
	return r._cat
}

// SetQ is Q Setter
// 商品筛选-查询词
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetQ(_q string) error {
	r._q = _q
	r.Set("q", _q)
	return nil
}

// GetQ Q Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetQ() string {
	return r._q
}

// SetIp is Ip Setter
// ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetIp() string {
	return r._ip
}

// SetDeviceencrypt is Deviceencrypt Setter
// 智能匹配-设备号加密类型：MD5
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetDeviceencrypt(_deviceencrypt string) error {
	r._deviceencrypt = _deviceencrypt
	r.Set("device_encrypt", _deviceencrypt)
	return nil
}

// GetDeviceencrypt Deviceencrypt Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetDeviceencrypt() string {
	return r._deviceencrypt
}

// SetDevicevalue is Devicevalue Setter
// 智能匹配-设备号加密后的值（MD5加密需32位小写）
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetDevicevalue(_devicevalue string) error {
	r._devicevalue = _devicevalue
	r.Set("device_value", _devicevalue)
	return nil
}

// GetDevicevalue Devicevalue Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetDevicevalue() string {
	return r._devicevalue
}

// SetDevicetype is Devicetype Setter
// 智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetDevicetype(_devicetype string) error {
	r._devicetype = _devicetype
	r.Set("device_type", _devicetype)
	return nil
}

// GetDevicetype Devicetype Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetDevicetype() string {
	return r._devicetype
}

// SetLongitude is Longitude Setter
// 本地化业务入参-LBS信息-经度
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetLatitude is Latitude Setter
// 本地化业务入参-LBS信息-纬度
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetCitycode is Citycode Setter
// 本地化业务入参-LBS信息-国标城市码，仅支持单个请求，请求饿了么卡券物料时，该字段必填。 （详细城市ID见：https://mo.m.taobao.com/page_2020010315120200508）
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetCitycode(_citycode string) error {
	r._citycode = _citycode
	r.Set("city_code", _citycode)
	return nil
}

// GetCitycode Citycode Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetCitycode() string {
	return r._citycode
}

// SetSellerids is Sellerids Setter
// 商家id，仅支持饿了么卡券商家ID，支持批量请求1-100以内，多个商家ID使用英文逗号分隔
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetSellerids(_sellerids string) error {
	r._sellerids = _sellerids
	r.Set("seller_ids", _sellerids)
	return nil
}

// GetSellerids Sellerids Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetSellerids() string {
	return r._sellerids
}

// SetSpecialid is Specialid Setter
// 会员运营ID
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetSpecialid(_specialid string) error {
	r._specialid = _specialid
	r.Set("special_id", _specialid)
	return nil
}

// GetSpecialid Specialid Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetSpecialid() string {
	return r._specialid
}

// SetRelationid is Relationid Setter
// 渠道关系ID，仅适用于渠道推广场景
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetRelationid(_relationid string) error {
	r._relationid = _relationid
	r.Set("relation_id", _relationid)
	return nil
}

// GetRelationid Relationid Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetRelationid() string {
	return r._relationid
}

// SetPageresultkey is Pageresultkey Setter
// 本地化业务入参-分页唯一标识，非首页的请求必传，值为上一页返回结果中的page_result_key字段值
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetPageresultkey(_pageresultkey string) error {
	r._pageresultkey = _pageresultkey
	r.Set("page_result_key", _pageresultkey)
	return nil
}

// GetPageresultkey Pageresultkey Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetPageresultkey() string {
	return r._pageresultkey
}

// SetBizsceneid is Bizsceneid Setter
// 1-动态ID转链场景，2-消费者比价场景（不填默认为1）
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetBizsceneid(_bizsceneid string) error {
	r._bizsceneid = _bizsceneid
	r.Set("biz_scene_id", _bizsceneid)
	return nil
}

// GetBizsceneid Bizsceneid Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetBizsceneid() string {
	return r._bizsceneid
}

// SetPromotiontype is Promotiontype Setter
// 1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetPromotiontype(_promotiontype string) error {
	r._promotiontype = _promotiontype
	r.Set("promotion_type", _promotiontype)
	return nil
}

// GetPromotiontype Promotiontype Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetPromotiontype() string {
	return r._promotiontype
}

// SetStartdsr is Startdsr Setter
// 商品筛选(特定媒体支持)-店铺dsr评分。筛选大于等于当前设置的店铺dsr评分的商品0-50000之间
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetStartdsr(_startdsr int64) error {
	r._startdsr = _startdsr
	r.Set("start_dsr", _startdsr)
	return nil
}

// GetStartdsr Startdsr Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetStartdsr() int64 {
	return r._startdsr
}

// SetPagesize is Pagesize Setter
// 页大小，默认20，1~100
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetPagesize(_pagesize int64) error {
	r._pagesize = _pagesize
	r.Set("page_size", _pagesize)
	return nil
}

// GetPagesize Pagesize Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetPagesize() int64 {
	return r._pagesize
}

// SetPageno is Pageno Setter
// 第几页，默认：１
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetPageno(_pageno int64) error {
	r._pageno = _pageno
	r.Set("page_no", _pageno)
	return nil
}

// GetPageno Pageno Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetPageno() int64 {
	return r._pageno
}

// SetPlatform is Platform Setter
// 链接形式：1：PC，2：无线，默认：１
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetEndtkrate is Endtkrate Setter
// 商品筛选-淘客佣金比率上限。如：1234表示12.34%
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetEndtkrate(_endtkrate int64) error {
	r._endtkrate = _endtkrate
	r.Set("end_tk_rate", _endtkrate)
	return nil
}

// GetEndtkrate Endtkrate Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetEndtkrate() int64 {
	return r._endtkrate
}

// SetStarttkrate is Starttkrate Setter
// 商品筛选-淘客佣金比率下限。如：1234表示12.34%
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetStarttkrate(_starttkrate int64) error {
	r._starttkrate = _starttkrate
	r.Set("start_tk_rate", _starttkrate)
	return nil
}

// GetStarttkrate Starttkrate Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetStarttkrate() int64 {
	return r._starttkrate
}

// SetEndprice is Endprice Setter
// 商品筛选-折扣价范围上限。单位：元
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetEndprice(_endprice int64) error {
	r._endprice = _endprice
	r.Set("end_price", _endprice)
	return nil
}

// GetEndprice Endprice Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetEndprice() int64 {
	return r._endprice
}

// SetStartprice is Startprice Setter
// 商品筛选-折扣价范围下限。单位：元
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetStartprice(_startprice int64) error {
	r._startprice = _startprice
	r.Set("start_price", _startprice)
	return nil
}

// GetStartprice Startprice Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetStartprice() int64 {
	return r._startprice
}

// SetMaterialid is Materialid Setter
// 不传时默认物料id=2836；如果直接对消费者投放，可使用官方个性化算法优化的搜索物料id=17004
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetMaterialid(_materialid int64) error {
	r._materialid = _materialid
	r.Set("material_id", _materialid)
	return nil
}

// GetMaterialid Materialid Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetMaterialid() int64 {
	return r._materialid
}

// SetAdzoneid is Adzoneid Setter
// mm_xxx_xxx_12345678三段式的最后一段数字
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetAdzoneid(_adzoneid int64) error {
	r._adzoneid = _adzoneid
	r.Set("adzone_id", _adzoneid)
	return nil
}

// GetAdzoneid Adzoneid Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetAdzoneid() int64 {
	return r._adzoneid
}

// SetNpxlevel is Npxlevel Setter
// 商品筛选-牛皮癣程度。取值：1不限，2无，3轻微
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetNpxlevel(_npxlevel int64) error {
	r._npxlevel = _npxlevel
	r.Set("npx_level", _npxlevel)
	return nil
}

// GetNpxlevel Npxlevel Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetNpxlevel() int64 {
	return r._npxlevel
}

// SetEndkatkrate is Endkatkrate Setter
// 商品筛选-KA媒体淘客佣金比率上限。如：1234表示12.34%
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetEndkatkrate(_endkatkrate int64) error {
	r._endkatkrate = _endkatkrate
	r.Set("end_ka_tk_rate", _endkatkrate)
	return nil
}

// GetEndkatkrate Endkatkrate Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetEndkatkrate() int64 {
	return r._endkatkrate
}

// SetStartkatkrate is Startkatkrate Setter
// 商品筛选-KA媒体淘客佣金比率下限。如：1234表示12.34%
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetStartkatkrate(_startkatkrate int64) error {
	r._startkatkrate = _startkatkrate
	r.Set("start_ka_tk_rate", _startkatkrate)
	return nil
}

// GetStartkatkrate Startkatkrate Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetStartkatkrate() int64 {
	return r._startkatkrate
}

// SetLockrateendtime is Lockrateendtime Setter
// 锁佣结束时间
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetLockrateendtime(_lockrateendtime int64) error {
	r._lockrateendtime = _lockrateendtime
	r.Set("lock_rate_end_time", _lockrateendtime)
	return nil
}

// GetLockrateendtime Lockrateendtime Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetLockrateendtime() int64 {
	return r._lockrateendtime
}

// SetLockratestarttime is Lockratestarttime Setter
// 锁佣开始时间
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetLockratestarttime(_lockratestarttime int64) error {
	r._lockratestarttime = _lockratestarttime
	r.Set("lock_rate_start_time", _lockratestarttime)
	return nil
}

// GetLockratestarttime Lockratestarttime Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetLockratestarttime() int64 {
	return r._lockratestarttime
}

// SetUcrowdid is Ucrowdid Setter
// 人群ID，仅适用于物料评估场景material_id=41377
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetUcrowdid(_ucrowdid int64) error {
	r._ucrowdid = _ucrowdid
	r.Set("ucrowd_id", _ucrowdid)
	return nil
}

// GetUcrowdid Ucrowdid Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetUcrowdid() int64 {
	return r._ucrowdid
}

// SetGettopnrate is Gettopnrate Setter
// 是否获取前N件佣金信息	0否，1是，其他值否
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetGettopnrate(_gettopnrate int64) error {
	r._gettopnrate = _gettopnrate
	r.Set("get_topn_rate", _gettopnrate)
	return nil
}

// GetGettopnrate Gettopnrate Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetGettopnrate() int64 {
	return r._gettopnrate
}

// SetIsoverseas is Isoverseas Setter
// 商品筛选-是否海外商品。true表示属于海外商品，false或不设置表示不限
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetIsoverseas(_isoverseas bool) error {
	r._isoverseas = _isoverseas
	r.Set("is_overseas", _isoverseas)
	return nil
}

// GetIsoverseas Isoverseas Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetIsoverseas() bool {
	return r._isoverseas
}

// SetIstmall is Istmall Setter
// 商品筛选-是否天猫商品。true表示属于天猫商品，false或不设置表示不限
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetIstmall(_istmall bool) error {
	r._istmall = _istmall
	r.Set("is_tmall", _istmall)
	return nil
}

// GetIstmall Istmall Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetIstmall() bool {
	return r._istmall
}

// SetHascoupon is Hascoupon Setter
// 优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetHascoupon(_hascoupon bool) error {
	r._hascoupon = _hascoupon
	r.Set("has_coupon", _hascoupon)
	return nil
}

// GetHascoupon Hascoupon Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetHascoupon() bool {
	return r._hascoupon
}

// SetNeedfreeshipment is Needfreeshipment Setter
// 商品筛选-是否包邮。true表示包邮，false或不设置表示不限
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetNeedfreeshipment(_needfreeshipment bool) error {
	r._needfreeshipment = _needfreeshipment
	r.Set("need_free_shipment", _needfreeshipment)
	return nil
}

// GetNeedfreeshipment Needfreeshipment Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetNeedfreeshipment() bool {
	return r._needfreeshipment
}

// SetNeedprepay is Needprepay Setter
// 商品筛选-是否加入消费者保障。true表示加入，false或不设置表示不限
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetNeedprepay(_needprepay bool) error {
	r._needprepay = _needprepay
	r.Set("need_prepay", _needprepay)
	return nil
}

// GetNeedprepay Needprepay Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetNeedprepay() bool {
	return r._needprepay
}

// SetIncludepayrate30 is Includepayrate30 Setter
// 商品筛选(特定媒体支持)-成交转化是否高于行业均值。True表示大于等于，false或不设置表示不限
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetIncludepayrate30(_includepayrate30 bool) error {
	r._includepayrate30 = _includepayrate30
	r.Set("include_pay_rate_30", _includepayrate30)
	return nil
}

// GetIncludepayrate30 Includepayrate30 Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetIncludepayrate30() bool {
	return r._includepayrate30
}

// SetIncludegoodrate is Includegoodrate Setter
// 商品筛选-好评率是否高于行业均值。True表示大于等于，false或不设置表示不限
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetIncludegoodrate(_includegoodrate bool) error {
	r._includegoodrate = _includegoodrate
	r.Set("include_good_rate", _includegoodrate)
	return nil
}

// GetIncludegoodrate Includegoodrate Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetIncludegoodrate() bool {
	return r._includegoodrate
}

// SetIncluderfdrate is Includerfdrate Setter
// 商品筛选(特定媒体支持)-退款率是否低于行业均值。True表示大于等于，false或不设置表示不限
func (r *TaobaotbkdgmaterialoptionalAPIRequest) SetIncluderfdrate(_includerfdrate bool) error {
	r._includerfdrate = _includerfdrate
	r.Set("include_rfd_rate", _includerfdrate)
	return nil
}

// GetIncluderfdrate Includerfdrate Getter
func (r TaobaotbkdgmaterialoptionalAPIRequest) GetIncluderfdrate() bool {
	return r._includerfdrate
}
