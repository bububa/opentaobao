package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
家装新零售主商品同步至阿里 API请求
tmall.nrt.item.main.synchronize

同步红星美凯龙存量商品到阿里
*/
type TmallNrtItemMainSynchronizeRequest struct {
    model.Params
    // 摊位id
    _boothId   string
    // 叶子类目id
    _cid   int64
    // 类目属性
    _props   []CategoryPropDto
    // 经销商编码
    _dealerCode   string
    // 卖场id
    _mallId   string
    // 商家编码
    _outerId   string
    // 系统自动生成
    _outerProps   *MacallineItemExtDto
    // 价格
    _price   string
    // 商品名
    _title   string
}

// 初始化TmallNrtItemMainSynchronizeRequest对象
func NewTmallNrtItemMainSynchronizeRequest() *TmallNrtItemMainSynchronizeRequest{
    return &TmallNrtItemMainSynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtItemMainSynchronizeRequest) GetApiMethodName() string {
    return "tmall.nrt.item.main.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtItemMainSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BoothId Setter
// 摊位id
func (r *TmallNrtItemMainSynchronizeRequest) SetBoothId(_boothId string) error {
    r._boothId = _boothId
    r.Set("booth_id", _boothId)
    return nil
}

// BoothId Getter
func (r TmallNrtItemMainSynchronizeRequest) GetBoothId() string {
    return r._boothId
}
// Cid Setter
// 叶子类目id
func (r *TmallNrtItemMainSynchronizeRequest) SetCid(_cid int64) error {
    r._cid = _cid
    r.Set("cid", _cid)
    return nil
}

// Cid Getter
func (r TmallNrtItemMainSynchronizeRequest) GetCid() int64 {
    return r._cid
}
// Props Setter
// 类目属性
func (r *TmallNrtItemMainSynchronizeRequest) SetProps(_props []CategoryPropDto) error {
    r._props = _props
    r.Set("props", _props)
    return nil
}

// Props Getter
func (r TmallNrtItemMainSynchronizeRequest) GetProps() []CategoryPropDto {
    return r._props
}
// DealerCode Setter
// 经销商编码
func (r *TmallNrtItemMainSynchronizeRequest) SetDealerCode(_dealerCode string) error {
    r._dealerCode = _dealerCode
    r.Set("dealer_code", _dealerCode)
    return nil
}

// DealerCode Getter
func (r TmallNrtItemMainSynchronizeRequest) GetDealerCode() string {
    return r._dealerCode
}
// MallId Setter
// 卖场id
func (r *TmallNrtItemMainSynchronizeRequest) SetMallId(_mallId string) error {
    r._mallId = _mallId
    r.Set("mall_id", _mallId)
    return nil
}

// MallId Getter
func (r TmallNrtItemMainSynchronizeRequest) GetMallId() string {
    return r._mallId
}
// OuterId Setter
// 商家编码
func (r *TmallNrtItemMainSynchronizeRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TmallNrtItemMainSynchronizeRequest) GetOuterId() string {
    return r._outerId
}
// OuterProps Setter
// 系统自动生成
func (r *TmallNrtItemMainSynchronizeRequest) SetOuterProps(_outerProps *MacallineItemExtDto) error {
    r._outerProps = _outerProps
    r.Set("outer_props", _outerProps)
    return nil
}

// OuterProps Getter
func (r TmallNrtItemMainSynchronizeRequest) GetOuterProps() *MacallineItemExtDto {
    return r._outerProps
}
// Price Setter
// 价格
func (r *TmallNrtItemMainSynchronizeRequest) SetPrice(_price string) error {
    r._price = _price
    r.Set("price", _price)
    return nil
}

// Price Getter
func (r TmallNrtItemMainSynchronizeRequest) GetPrice() string {
    return r._price
}
// Title Setter
// 商品名
func (r *TmallNrtItemMainSynchronizeRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TmallNrtItemMainSynchronizeRequest) GetTitle() string {
    return r._title
}
