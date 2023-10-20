package openmall

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeAddressParseAPIRequest openmall服务地址区域码解析 API请求
// taobao.openmall.trade.address.parse
//
// openmall服务，解析地址区域码，获取创建订单等接口中的区域码信息
type TaobaoOpenmallTradeAddressParseAPIRequest struct {
	model.Params
	// 需解析的地址信息，建议只传地址选择器中的省市区，街道门牌号等用户手动输入数据不传
	_rawAddress string
	// 渠道商分销者淘宝账号
	_distributor string
}

// NewTaobaoOpenmallTradeAddressParseRequest 初始化TaobaoOpenmallTradeAddressParseAPIRequest对象
func NewTaobaoOpenmallTradeAddressParseRequest() *TaobaoOpenmallTradeAddressParseAPIRequest {
	return &TaobaoOpenmallTradeAddressParseAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenmallTradeAddressParseAPIRequest) Reset() {
	r._rawAddress = ""
	r._distributor = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeAddressParseAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.address.parse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeAddressParseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallTradeAddressParseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRawAddress is RawAddress Setter
// 需解析的地址信息，建议只传地址选择器中的省市区，街道门牌号等用户手动输入数据不传
func (r *TaobaoOpenmallTradeAddressParseAPIRequest) SetRawAddress(_rawAddress string) error {
	r._rawAddress = _rawAddress
	r.Set("raw_address", _rawAddress)
	return nil
}

// GetRawAddress RawAddress Getter
func (r TaobaoOpenmallTradeAddressParseAPIRequest) GetRawAddress() string {
	return r._rawAddress
}

// SetDistributor is Distributor Setter
// 渠道商分销者淘宝账号
func (r *TaobaoOpenmallTradeAddressParseAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallTradeAddressParseAPIRequest) GetDistributor() string {
	return r._distributor
}

var poolTaobaoOpenmallTradeAddressParseAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenmallTradeAddressParseRequest()
	},
}

// GetTaobaoOpenmallTradeAddressParseRequest 从 sync.Pool 获取 TaobaoOpenmallTradeAddressParseAPIRequest
func GetTaobaoOpenmallTradeAddressParseAPIRequest() *TaobaoOpenmallTradeAddressParseAPIRequest {
	return poolTaobaoOpenmallTradeAddressParseAPIRequest.Get().(*TaobaoOpenmallTradeAddressParseAPIRequest)
}

// ReleaseTaobaoOpenmallTradeAddressParseAPIRequest 将 TaobaoOpenmallTradeAddressParseAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenmallTradeAddressParseAPIRequest(v *TaobaoOpenmallTradeAddressParseAPIRequest) {
	v.Reset()
	poolTaobaoOpenmallTradeAddressParseAPIRequest.Put(v)
}
