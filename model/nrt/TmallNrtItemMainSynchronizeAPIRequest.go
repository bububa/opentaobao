package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtItemMainSynchronizeAPIRequest 家装新零售主商品同步至阿里 API请求
// tmall.nrt.item.main.synchronize
//
// 同步红星美凯龙存量商品到阿里
type TmallNrtItemMainSynchronizeAPIRequest struct {
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

// NewTmallNrtItemMainSynchronizeRequest 初始化TmallNrtItemMainSynchronizeAPIRequest对象
func NewTmallNrtItemMainSynchronizeRequest() *TmallNrtItemMainSynchronizeAPIRequest {
	return &TmallNrtItemMainSynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtItemMainSynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.item.main.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtItemMainSynchronizeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetProps is Props Setter
// 类目属性
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetProps(_props []CategoryPropDto) error {
	r._props = _props
	r.Set("props", _props)
	return nil
}

// GetProps Props Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetProps() []CategoryPropDto {
	return r._props
}

// SetBoothId is BoothId Setter
// 摊位id
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetBoothId(_boothId string) error {
	r._boothId = _boothId
	r.Set("booth_id", _boothId)
	return nil
}

// GetBoothId BoothId Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetBoothId() string {
	return r._boothId
}

// SetDealerCode is DealerCode Setter
// 经销商编码
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetDealerCode(_dealerCode string) error {
	r._dealerCode = _dealerCode
	r.Set("dealer_code", _dealerCode)
	return nil
}

// GetDealerCode DealerCode Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetDealerCode() string {
	return r._dealerCode
}

// SetMallId is MallId Setter
// 卖场id
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetMallId(_mallId string) error {
	r._mallId = _mallId
	r.Set("mall_id", _mallId)
	return nil
}

// GetMallId MallId Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetMallId() string {
	return r._mallId
}

// SetOuterId is OuterId Setter
// 商家编码
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetPrice is Price Setter
// 价格
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetPrice(_price string) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetPrice() string {
	return r._price
}

// SetTitle is Title Setter
// 商品名
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetTitle() string {
	return r._title
}

// SetCid is Cid Setter
// 叶子类目id
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// GetCid Cid Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetCid() int64 {
	return r._cid
}

// SetOuterProps is OuterProps Setter
// 系统自动生成
func (r *TmallNrtItemMainSynchronizeAPIRequest) SetOuterProps(_outerProps *MacallineItemExtDto) error {
	r._outerProps = _outerProps
	r.Set("outer_props", _outerProps)
	return nil
}

// GetOuterProps OuterProps Getter
func (r TmallNrtItemMainSynchronizeAPIRequest) GetOuterProps() *MacallineItemExtDto {
	return r._outerProps
}
