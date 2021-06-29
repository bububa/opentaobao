package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-店铺搜索 API请求
taobao.tbk.shop.get

淘宝客店铺查询
*/
type TaobaoTbkShopGetRequest struct {
    model.Params
    // 需返回的字段列表
    _fields   string
    // 查询词
    _q   string
    // 排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction）
    _sort   string
    // 是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性
    _isTmall   bool
    // 信用等级下限，1~20
    _startCredit   int64
    // 信用等级上限，1~20
    _endCredit   int64
    // 淘客佣金比率下限，1~10000
    _startCommissionRate   int64
    // 淘客佣金比率上限，1~10000
    _endCommissionRate   int64
    // 店铺商品总数下限
    _startTotalAction   int64
    // 店铺商品总数上限
    _endTotalAction   int64
    // 累计推广商品下限
    _startAuctionCount   int64
    // 累计推广商品上限
    _endAuctionCount   int64
    // 链接形式：1：PC，2：无线，默认：１
    _platform   int64
    // 第几页，默认1，1~100
    _pageNo   int64
    // 页大小，默认20，1~100
    _pageSize   int64
}

// 初始化TaobaoTbkShopGetRequest对象
func NewTaobaoTbkShopGetRequest() *TaobaoTbkShopGetRequest{
    return &TaobaoTbkShopGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkShopGetRequest) GetApiMethodName() string {
    return "taobao.tbk.shop.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkShopGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需返回的字段列表
func (r *TaobaoTbkShopGetRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoTbkShopGetRequest) GetFields() string {
    return r._fields
}
// Q Setter
// 查询词
func (r *TaobaoTbkShopGetRequest) SetQ(_q string) error {
    r._q = _q
    r.Set("q", _q)
    return nil
}

// Q Getter
func (r TaobaoTbkShopGetRequest) GetQ() string {
    return r._q
}
// Sort Setter
// 排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction）
func (r *TaobaoTbkShopGetRequest) SetSort(_sort string) error {
    r._sort = _sort
    r.Set("sort", _sort)
    return nil
}

// Sort Getter
func (r TaobaoTbkShopGetRequest) GetSort() string {
    return r._sort
}
// IsTmall Setter
// 是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性
func (r *TaobaoTbkShopGetRequest) SetIsTmall(_isTmall bool) error {
    r._isTmall = _isTmall
    r.Set("is_tmall", _isTmall)
    return nil
}

// IsTmall Getter
func (r TaobaoTbkShopGetRequest) GetIsTmall() bool {
    return r._isTmall
}
// StartCredit Setter
// 信用等级下限，1~20
func (r *TaobaoTbkShopGetRequest) SetStartCredit(_startCredit int64) error {
    r._startCredit = _startCredit
    r.Set("start_credit", _startCredit)
    return nil
}

// StartCredit Getter
func (r TaobaoTbkShopGetRequest) GetStartCredit() int64 {
    return r._startCredit
}
// EndCredit Setter
// 信用等级上限，1~20
func (r *TaobaoTbkShopGetRequest) SetEndCredit(_endCredit int64) error {
    r._endCredit = _endCredit
    r.Set("end_credit", _endCredit)
    return nil
}

// EndCredit Getter
func (r TaobaoTbkShopGetRequest) GetEndCredit() int64 {
    return r._endCredit
}
// StartCommissionRate Setter
// 淘客佣金比率下限，1~10000
func (r *TaobaoTbkShopGetRequest) SetStartCommissionRate(_startCommissionRate int64) error {
    r._startCommissionRate = _startCommissionRate
    r.Set("start_commission_rate", _startCommissionRate)
    return nil
}

// StartCommissionRate Getter
func (r TaobaoTbkShopGetRequest) GetStartCommissionRate() int64 {
    return r._startCommissionRate
}
// EndCommissionRate Setter
// 淘客佣金比率上限，1~10000
func (r *TaobaoTbkShopGetRequest) SetEndCommissionRate(_endCommissionRate int64) error {
    r._endCommissionRate = _endCommissionRate
    r.Set("end_commission_rate", _endCommissionRate)
    return nil
}

// EndCommissionRate Getter
func (r TaobaoTbkShopGetRequest) GetEndCommissionRate() int64 {
    return r._endCommissionRate
}
// StartTotalAction Setter
// 店铺商品总数下限
func (r *TaobaoTbkShopGetRequest) SetStartTotalAction(_startTotalAction int64) error {
    r._startTotalAction = _startTotalAction
    r.Set("start_total_action", _startTotalAction)
    return nil
}

// StartTotalAction Getter
func (r TaobaoTbkShopGetRequest) GetStartTotalAction() int64 {
    return r._startTotalAction
}
// EndTotalAction Setter
// 店铺商品总数上限
func (r *TaobaoTbkShopGetRequest) SetEndTotalAction(_endTotalAction int64) error {
    r._endTotalAction = _endTotalAction
    r.Set("end_total_action", _endTotalAction)
    return nil
}

// EndTotalAction Getter
func (r TaobaoTbkShopGetRequest) GetEndTotalAction() int64 {
    return r._endTotalAction
}
// StartAuctionCount Setter
// 累计推广商品下限
func (r *TaobaoTbkShopGetRequest) SetStartAuctionCount(_startAuctionCount int64) error {
    r._startAuctionCount = _startAuctionCount
    r.Set("start_auction_count", _startAuctionCount)
    return nil
}

// StartAuctionCount Getter
func (r TaobaoTbkShopGetRequest) GetStartAuctionCount() int64 {
    return r._startAuctionCount
}
// EndAuctionCount Setter
// 累计推广商品上限
func (r *TaobaoTbkShopGetRequest) SetEndAuctionCount(_endAuctionCount int64) error {
    r._endAuctionCount = _endAuctionCount
    r.Set("end_auction_count", _endAuctionCount)
    return nil
}

// EndAuctionCount Getter
func (r TaobaoTbkShopGetRequest) GetEndAuctionCount() int64 {
    return r._endAuctionCount
}
// Platform Setter
// 链接形式：1：PC，2：无线，默认：１
func (r *TaobaoTbkShopGetRequest) SetPlatform(_platform int64) error {
    r._platform = _platform
    r.Set("platform", _platform)
    return nil
}

// Platform Getter
func (r TaobaoTbkShopGetRequest) GetPlatform() int64 {
    return r._platform
}
// PageNo Setter
// 第几页，默认1，1~100
func (r *TaobaoTbkShopGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTbkShopGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaoTbkShopGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTbkShopGetRequest) GetPageSize() int64 {
    return r._pageSize
}
