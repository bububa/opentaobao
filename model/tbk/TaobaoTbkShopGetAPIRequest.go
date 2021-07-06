package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkShopGetAPIRequest 淘宝客-推广者-店铺搜索 API请求
// taobao.tbk.shop.get
//
// 淘宝客店铺查询
type TaobaoTbkShopGetAPIRequest struct {
	model.Params
	// 需返回的字段列表
	_fields string
	// 查询词
	_q string
	// 排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction）
	_sort string
	// 信用等级下限，1~20
	_startCredit int64
	// 信用等级上限，1~20
	_endCredit int64
	// 淘客佣金比率下限，1~10000
	_startCommissionRate int64
	// 淘客佣金比率上限，1~10000
	_endCommissionRate int64
	// 店铺商品总数下限
	_startTotalAction int64
	// 店铺商品总数上限
	_endTotalAction int64
	// 累计推广商品下限
	_startAuctionCount int64
	// 累计推广商品上限
	_endAuctionCount int64
	// 链接形式：1：PC，2：无线，默认：１
	_platform int64
	// 第几页，默认1，1~100
	_pageNo int64
	// 页大小，默认20，1~100
	_pageSize int64
	// 是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性
	_isTmall bool
}

// NewTaobaoTbkShopGetRequest 初始化TaobaoTbkShopGetAPIRequest对象
func NewTaobaoTbkShopGetRequest() *TaobaoTbkShopGetAPIRequest {
	return &TaobaoTbkShopGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkShopGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.shop.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkShopGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 需返回的字段列表
func (r *TaobaoTbkShopGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTbkShopGetAPIRequest) GetFields() string {
	return r._fields
}

// SetQ is Q Setter
// 查询词
func (r *TaobaoTbkShopGetAPIRequest) SetQ(_q string) error {
	r._q = _q
	r.Set("q", _q)
	return nil
}

// GetQ Q Getter
func (r TaobaoTbkShopGetAPIRequest) GetQ() string {
	return r._q
}

// SetSort is Sort Setter
// 排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction）
func (r *TaobaoTbkShopGetAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r TaobaoTbkShopGetAPIRequest) GetSort() string {
	return r._sort
}

// SetStartCredit is StartCredit Setter
// 信用等级下限，1~20
func (r *TaobaoTbkShopGetAPIRequest) SetStartCredit(_startCredit int64) error {
	r._startCredit = _startCredit
	r.Set("start_credit", _startCredit)
	return nil
}

// GetStartCredit StartCredit Getter
func (r TaobaoTbkShopGetAPIRequest) GetStartCredit() int64 {
	return r._startCredit
}

// SetEndCredit is EndCredit Setter
// 信用等级上限，1~20
func (r *TaobaoTbkShopGetAPIRequest) SetEndCredit(_endCredit int64) error {
	r._endCredit = _endCredit
	r.Set("end_credit", _endCredit)
	return nil
}

// GetEndCredit EndCredit Getter
func (r TaobaoTbkShopGetAPIRequest) GetEndCredit() int64 {
	return r._endCredit
}

// SetStartCommissionRate is StartCommissionRate Setter
// 淘客佣金比率下限，1~10000
func (r *TaobaoTbkShopGetAPIRequest) SetStartCommissionRate(_startCommissionRate int64) error {
	r._startCommissionRate = _startCommissionRate
	r.Set("start_commission_rate", _startCommissionRate)
	return nil
}

// GetStartCommissionRate StartCommissionRate Getter
func (r TaobaoTbkShopGetAPIRequest) GetStartCommissionRate() int64 {
	return r._startCommissionRate
}

// SetEndCommissionRate is EndCommissionRate Setter
// 淘客佣金比率上限，1~10000
func (r *TaobaoTbkShopGetAPIRequest) SetEndCommissionRate(_endCommissionRate int64) error {
	r._endCommissionRate = _endCommissionRate
	r.Set("end_commission_rate", _endCommissionRate)
	return nil
}

// GetEndCommissionRate EndCommissionRate Getter
func (r TaobaoTbkShopGetAPIRequest) GetEndCommissionRate() int64 {
	return r._endCommissionRate
}

// SetStartTotalAction is StartTotalAction Setter
// 店铺商品总数下限
func (r *TaobaoTbkShopGetAPIRequest) SetStartTotalAction(_startTotalAction int64) error {
	r._startTotalAction = _startTotalAction
	r.Set("start_total_action", _startTotalAction)
	return nil
}

// GetStartTotalAction StartTotalAction Getter
func (r TaobaoTbkShopGetAPIRequest) GetStartTotalAction() int64 {
	return r._startTotalAction
}

// SetEndTotalAction is EndTotalAction Setter
// 店铺商品总数上限
func (r *TaobaoTbkShopGetAPIRequest) SetEndTotalAction(_endTotalAction int64) error {
	r._endTotalAction = _endTotalAction
	r.Set("end_total_action", _endTotalAction)
	return nil
}

// GetEndTotalAction EndTotalAction Getter
func (r TaobaoTbkShopGetAPIRequest) GetEndTotalAction() int64 {
	return r._endTotalAction
}

// SetStartAuctionCount is StartAuctionCount Setter
// 累计推广商品下限
func (r *TaobaoTbkShopGetAPIRequest) SetStartAuctionCount(_startAuctionCount int64) error {
	r._startAuctionCount = _startAuctionCount
	r.Set("start_auction_count", _startAuctionCount)
	return nil
}

// GetStartAuctionCount StartAuctionCount Getter
func (r TaobaoTbkShopGetAPIRequest) GetStartAuctionCount() int64 {
	return r._startAuctionCount
}

// SetEndAuctionCount is EndAuctionCount Setter
// 累计推广商品上限
func (r *TaobaoTbkShopGetAPIRequest) SetEndAuctionCount(_endAuctionCount int64) error {
	r._endAuctionCount = _endAuctionCount
	r.Set("end_auction_count", _endAuctionCount)
	return nil
}

// GetEndAuctionCount EndAuctionCount Getter
func (r TaobaoTbkShopGetAPIRequest) GetEndAuctionCount() int64 {
	return r._endAuctionCount
}

// SetPlatform is Platform Setter
// 链接形式：1：PC，2：无线，默认：１
func (r *TaobaoTbkShopGetAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaoTbkShopGetAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetPageNo is PageNo Setter
// 第几页，默认1，1~100
func (r *TaobaoTbkShopGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoTbkShopGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaoTbkShopGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTbkShopGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetIsTmall is IsTmall Setter
// 是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性
func (r *TaobaoTbkShopGetAPIRequest) SetIsTmall(_isTmall bool) error {
	r._isTmall = _isTmall
	r.Set("is_tmall", _isTmall)
	return nil
}

// GetIsTmall IsTmall Getter
func (r TaobaoTbkShopGetAPIRequest) GetIsTmall() bool {
	return r._isTmall
}
