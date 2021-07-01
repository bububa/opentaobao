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
type TmallNrtItemMainSynchronizeAPIRequest struct {
    model.Params
    // 摊位id
    _boothId   string
    // 叶子类目id
    _cid   int64
    // 类目属性
    _props   []CategoryPropDTO
    // 经销商编码
    _dealerCode   string
    // 卖场id
    _mallId   string
    // 商家编码
    _outerId   string
    // 系统自动生成
    _outerProps   *MacallineItemExtDTO
    // 价格
    _price   string
    // 商品名
    _title   string
}

// 初始化TmallNrtItemMainSynchronizeAPIRequest对象
func NewTmallNrtItemMainSynchronizeRequest() *TmallNrtItemMainSynchronizeAPIRequest{
    return &TmallNrtItemMainSynchronizeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtItemMainSynchronizeAPIRequest) GetApiMethodName() string {
    return "tmall.nrt.item.main.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtItemMainSynchronizeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BoothId Setter
// 摊位id
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetBoothId(_boothId string) error {
    r._boothId = _boothId
    r.Set("booth_id", _boothId)
    return nil
}

// BoothId Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetBoothId() string {
    return r._boothId
}
// Cid Setter
// 叶子类目id
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetCid(_cid int64) error {
    r._cid = _cid
    r.Set("cid", _cid)
    return nil
}

// Cid Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetCid() int64 {
    return r._cid
}
// Props Setter
// 类目属性
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetProps(_props []CategoryPropDTO) error {
    r._props = _props
    r.Set("props", _props)
    return nil
}

// Props Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetProps() []CategoryPropDTO {
    return r._props
}
// DealerCode Setter
// 经销商编码
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetDealerCode(_dealerCode string) error {
    r._dealerCode = _dealerCode
    r.Set("dealer_code", _dealerCode)
    return nil
}

// DealerCode Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetDealerCode() string {
    return r._dealerCode
}
// MallId Setter
// 卖场id
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetMallId(_mallId string) error {
    r._mallId = _mallId
    r.Set("mall_id", _mallId)
    return nil
}

// MallId Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetMallId() string {
    return r._mallId
}
// OuterId Setter
// 商家编码
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetOuterId() string {
    return r._outerId
}
// OuterProps Setter
// 系统自动生成
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetOuterProps(_outerProps *MacallineItemExtDTO) error {
    r._outerProps = _outerProps
    r.Set("outer_props", _outerProps)
    return nil
}

// OuterProps Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetOuterProps() *MacallineItemExtDTO {
    return r._outerProps
}
// Price Setter
// 价格
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetPrice(_price string) error {
    r._price = _price
    r.Set("price", _price)
    return nil
}

// Price Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetPrice() string {
    return r._price
}
// Title Setter
// 商品名
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetTitle() string {
    return r._title
}
