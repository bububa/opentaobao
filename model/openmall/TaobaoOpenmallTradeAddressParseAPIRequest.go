package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmalltradeaddressparseAPIRequest openmall服务地址区域码解析 API请求
// taobao.openmall.trade.address.parse
//
// openmall服务，解析地址区域码，获取创建订单等接口中的区域码信息
type TaobaoopenmalltradeaddressparseAPIRequest struct {
	model.Params
	// 需解析的地址信息，建议只传地址选择器中的省市区，街道门牌号等用户手动输入数据不传
	_rawAddress string
	// 渠道商分销者淘宝账号
	_distributor string
}

// NewTaobaoopenmalltradeaddressparseRequest 初始化TaobaoopenmalltradeaddressparseAPIRequest对象
func NewTaobaoopenmalltradeaddressparseRequest() *TaobaoopenmalltradeaddressparseAPIRequest {
	return &TaobaoopenmalltradeaddressparseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmalltradeaddressparseAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.address.parse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmalltradeaddressparseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmalltradeaddressparseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRawAddress is RawAddress Setter
// 需解析的地址信息，建议只传地址选择器中的省市区，街道门牌号等用户手动输入数据不传
func (r *TaobaoopenmalltradeaddressparseAPIRequest) SetRawAddress(_rawAddress string) error {
	r._rawAddress = _rawAddress
	r.Set("raw_address", _rawAddress)
	return nil
}

// GetRawAddress RawAddress Getter
func (r TaobaoopenmalltradeaddressparseAPIRequest) GetRawAddress() string {
	return r._rawAddress
}

// SetDistributor is Distributor Setter
// 渠道商分销者淘宝账号
func (r *TaobaoopenmalltradeaddressparseAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoopenmalltradeaddressparseAPIRequest) GetDistributor() string {
	return r._distributor
}
