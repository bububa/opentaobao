package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtitemmainsynchronizeAPIRequest 家装新零售主商品同步至阿里 API请求
// tmall.nrt.item.main.synchronize
//
// 同步卖场存量线下商品到阿里
type TmallnrtitemmainsynchronizeAPIRequest struct {
	model.Params
	// 类目属性
	_props []CategoryPropDto
	// 摊位id
	_boothId string
	// 经销商编码
	_dealerCode string
	// 卖场id
	_mallId string
	// 商家编码
	_outerId string
	// 价格
	_price string
	// 商品名
	_title string
	// 叶子类目id
	_cid int64
	// 系统自动生成
	_outerProps *MacallineItemExtDto
}

// NewTmallnrtitemmainsynchronizeRequest 初始化TmallnrtitemmainsynchronizeAPIRequest对象
func NewTmallnrtitemmainsynchronizeRequest() *TmallnrtitemmainsynchronizeAPIRequest {
	return &TmallnrtitemmainsynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtitemmainsynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.item.main.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtitemmainsynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtitemmainsynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProps is Props Setter
// 类目属性
func (r *TmallnrtitemmainsynchronizeAPIRequest) SetProps(_props []CategoryPropDto) error {
	r._props = _props
	r.Set("props", _props)
	return nil
}

// GetProps Props Getter
func (r TmallnrtitemmainsynchronizeAPIRequest) GetProps() []CategoryPropDto {
	return r._props
}

// SetBoothId is BoothId Setter
// 摊位id
func (r *TmallnrtitemmainsynchronizeAPIRequest) SetBoothId(_boothId string) error {
	r._boothId = _boothId
	r.Set("booth_id", _boothId)
	return nil
}

// GetBoothId BoothId Getter
func (r TmallnrtitemmainsynchronizeAPIRequest) GetBoothId() string {
	return r._boothId
}

// SetDealerCode is DealerCode Setter
// 经销商编码
func (r *TmallnrtitemmainsynchronizeAPIRequest) SetDealerCode(_dealerCode string) error {
	r._dealerCode = _dealerCode
	r.Set("dealer_code", _dealerCode)
	return nil
}

// GetDealerCode DealerCode Getter
func (r TmallnrtitemmainsynchronizeAPIRequest) GetDealerCode() string {
	return r._dealerCode
}

// SetMallId is MallId Setter
// 卖场id
func (r *TmallnrtitemmainsynchronizeAPIRequest) SetMallId(_mallId string) error {
	r._mallId = _mallId
	r.Set("mall_id", _mallId)
	return nil
}

// GetMallId MallId Getter
func (r TmallnrtitemmainsynchronizeAPIRequest) GetMallId() string {
	return r._mallId
}

// SetOuterId is OuterId Setter
// 商家编码
func (r *TmallnrtitemmainsynchronizeAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TmallnrtitemmainsynchronizeAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetPrice is Price Setter
// 价格
func (r *TmallnrtitemmainsynchronizeAPIRequest) SetPrice(_price string) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r TmallnrtitemmainsynchronizeAPIRequest) GetPrice() string {
	return r._price
}

// SetTitle is Title Setter
// 商品名
func (r *TmallnrtitemmainsynchronizeAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TmallnrtitemmainsynchronizeAPIRequest) GetTitle() string {
	return r._title
}

// SetCid is Cid Setter
// 叶子类目id
func (r *TmallnrtitemmainsynchronizeAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// GetCid Cid Getter
func (r TmallnrtitemmainsynchronizeAPIRequest) GetCid() int64 {
	return r._cid
}

// SetOuterProps is OuterProps Setter
// 系统自动生成
func (r *TmallnrtitemmainsynchronizeAPIRequest) SetOuterProps(_outerProps *MacallineItemExtDto) error {
	r._outerProps = _outerProps
	r.Set("outer_props", _outerProps)
	return nil
}

// GetOuterProps OuterProps Getter
func (r TmallnrtitemmainsynchronizeAPIRequest) GetOuterProps() *MacallineItemExtDto {
	return r._outerProps
}
