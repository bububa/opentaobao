package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkshopgetAPIRequest 淘宝客-推广者-店铺搜索 API请求
// taobao.tbk.shop.get
//
// 淘宝客店铺查询
type TaobaotbkshopgetAPIRequest struct {
	model.Params
	// 需返回的字段列表
	_fields string
	// 查询词
	_q string
	// 排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction）
	_sort string
	// 累计推广商品上限
	_endauctioncount int64
	// 淘客佣金比率上限，1~10000
	_endcommissionrate int64
	// 信用等级上限，1~20
	_endcredit int64
	// 店铺商品总数上限
	_endtotalaction int64
	// 第几页，默认1，1~100
	_pageno int64
	// 页大小，默认20，1~100
	_pagesize int64
	// 链接形式：1：PC，2：无线，默认：１
	_platform int64
	// 累计推广商品下限
	_startauctioncount int64
	// 淘客佣金比率下限，1~10000
	_startcommissionrate int64
	// 信用等级下限，1~20
	_startcredit int64
	// 店铺商品总数下限
	_starttotalaction int64
	// 是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性
	_istmall bool
}

// NewTaobaotbkshopgetRequest 初始化TaobaotbkshopgetAPIRequest对象
func NewTaobaotbkshopgetRequest() *TaobaotbkshopgetAPIRequest {
	return &TaobaotbkshopgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkshopgetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.shop.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkshopgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkshopgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表
func (r *TaobaotbkshopgetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaotbkshopgetAPIRequest) GetFields() string {
	return r._fields
}

// SetQ is Q Setter
// 查询词
func (r *TaobaotbkshopgetAPIRequest) SetQ(_q string) error {
	r._q = _q
	r.Set("q", _q)
	return nil
}

// GetQ Q Getter
func (r TaobaotbkshopgetAPIRequest) GetQ() string {
	return r._q
}

// SetSort is Sort Setter
// 排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction）
func (r *TaobaotbkshopgetAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r TaobaotbkshopgetAPIRequest) GetSort() string {
	return r._sort
}

// SetEndauctioncount is Endauctioncount Setter
// 累计推广商品上限
func (r *TaobaotbkshopgetAPIRequest) SetEndauctioncount(_endauctioncount int64) error {
	r._endauctioncount = _endauctioncount
	r.Set("end_auction_count", _endauctioncount)
	return nil
}

// GetEndauctioncount Endauctioncount Getter
func (r TaobaotbkshopgetAPIRequest) GetEndauctioncount() int64 {
	return r._endauctioncount
}

// SetEndcommissionrate is Endcommissionrate Setter
// 淘客佣金比率上限，1~10000
func (r *TaobaotbkshopgetAPIRequest) SetEndcommissionrate(_endcommissionrate int64) error {
	r._endcommissionrate = _endcommissionrate
	r.Set("end_commission_rate", _endcommissionrate)
	return nil
}

// GetEndcommissionrate Endcommissionrate Getter
func (r TaobaotbkshopgetAPIRequest) GetEndcommissionrate() int64 {
	return r._endcommissionrate
}

// SetEndcredit is Endcredit Setter
// 信用等级上限，1~20
func (r *TaobaotbkshopgetAPIRequest) SetEndcredit(_endcredit int64) error {
	r._endcredit = _endcredit
	r.Set("end_credit", _endcredit)
	return nil
}

// GetEndcredit Endcredit Getter
func (r TaobaotbkshopgetAPIRequest) GetEndcredit() int64 {
	return r._endcredit
}

// SetEndtotalaction is Endtotalaction Setter
// 店铺商品总数上限
func (r *TaobaotbkshopgetAPIRequest) SetEndtotalaction(_endtotalaction int64) error {
	r._endtotalaction = _endtotalaction
	r.Set("end_total_action", _endtotalaction)
	return nil
}

// GetEndtotalaction Endtotalaction Getter
func (r TaobaotbkshopgetAPIRequest) GetEndtotalaction() int64 {
	return r._endtotalaction
}

// SetPageno is Pageno Setter
// 第几页，默认1，1~100
func (r *TaobaotbkshopgetAPIRequest) SetPageno(_pageno int64) error {
	r._pageno = _pageno
	r.Set("page_no", _pageno)
	return nil
}

// GetPageno Pageno Getter
func (r TaobaotbkshopgetAPIRequest) GetPageno() int64 {
	return r._pageno
}

// SetPagesize is Pagesize Setter
// 页大小，默认20，1~100
func (r *TaobaotbkshopgetAPIRequest) SetPagesize(_pagesize int64) error {
	r._pagesize = _pagesize
	r.Set("page_size", _pagesize)
	return nil
}

// GetPagesize Pagesize Getter
func (r TaobaotbkshopgetAPIRequest) GetPagesize() int64 {
	return r._pagesize
}

// SetPlatform is Platform Setter
// 链接形式：1：PC，2：无线，默认：１
func (r *TaobaotbkshopgetAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaotbkshopgetAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetStartauctioncount is Startauctioncount Setter
// 累计推广商品下限
func (r *TaobaotbkshopgetAPIRequest) SetStartauctioncount(_startauctioncount int64) error {
	r._startauctioncount = _startauctioncount
	r.Set("start_auction_count", _startauctioncount)
	return nil
}

// GetStartauctioncount Startauctioncount Getter
func (r TaobaotbkshopgetAPIRequest) GetStartauctioncount() int64 {
	return r._startauctioncount
}

// SetStartcommissionrate is Startcommissionrate Setter
// 淘客佣金比率下限，1~10000
func (r *TaobaotbkshopgetAPIRequest) SetStartcommissionrate(_startcommissionrate int64) error {
	r._startcommissionrate = _startcommissionrate
	r.Set("start_commission_rate", _startcommissionrate)
	return nil
}

// GetStartcommissionrate Startcommissionrate Getter
func (r TaobaotbkshopgetAPIRequest) GetStartcommissionrate() int64 {
	return r._startcommissionrate
}

// SetStartcredit is Startcredit Setter
// 信用等级下限，1~20
func (r *TaobaotbkshopgetAPIRequest) SetStartcredit(_startcredit int64) error {
	r._startcredit = _startcredit
	r.Set("start_credit", _startcredit)
	return nil
}

// GetStartcredit Startcredit Getter
func (r TaobaotbkshopgetAPIRequest) GetStartcredit() int64 {
	return r._startcredit
}

// SetStarttotalaction is Starttotalaction Setter
// 店铺商品总数下限
func (r *TaobaotbkshopgetAPIRequest) SetStarttotalaction(_starttotalaction int64) error {
	r._starttotalaction = _starttotalaction
	r.Set("start_total_action", _starttotalaction)
	return nil
}

// GetStarttotalaction Starttotalaction Getter
func (r TaobaotbkshopgetAPIRequest) GetStarttotalaction() int64 {
	return r._starttotalaction
}

// SetIstmall is Istmall Setter
// 是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性
func (r *TaobaotbkshopgetAPIRequest) SetIstmall(_istmall bool) error {
	r._istmall = _istmall
	r.Set("is_tmall", _istmall)
	return nil
}

// GetIstmall Istmall Getter
func (r TaobaotbkshopgetAPIRequest) GetIstmall() bool {
	return r._istmall
}
