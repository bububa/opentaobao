package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest 阿信酒店分销-订单详情接口 API请求
// taobao.alitrip.travel.axin.hotel.order.detail
//
// 阿信酒店订单详情的读取接口
type TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest struct {
	model.Params
	// 外部订单id
	_outerOrderId string
	// 资源渠道
	_resourceChannel string
	// 采购单id
	_purSubOrderId int64
	// 分销商id
	_distributorTid int64
}

// NewTaobaoAlitripTravelAxinHotelOrderDetailRequest 初始化TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelOrderDetailRequest() *TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest {
	return &TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.order.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterOrderId is OuterOrderId Setter
// 外部订单id
func (r *TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest) SetOuterOrderId(_outerOrderId string) error {
	r._outerOrderId = _outerOrderId
	r.Set("outer_order_id", _outerOrderId)
	return nil
}

// GetOuterOrderId OuterOrderId Getter
func (r TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest) GetOuterOrderId() string {
	return r._outerOrderId
}

// SetResourceChannel is ResourceChannel Setter
// 资源渠道
func (r *TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest) SetResourceChannel(_resourceChannel string) error {
	r._resourceChannel = _resourceChannel
	r.Set("resource_channel", _resourceChannel)
	return nil
}

// GetResourceChannel ResourceChannel Getter
func (r TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest) GetResourceChannel() string {
	return r._resourceChannel
}

// SetPurSubOrderId is PurSubOrderId Setter
// 采购单id
func (r *TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest) SetPurSubOrderId(_purSubOrderId int64) error {
	r._purSubOrderId = _purSubOrderId
	r.Set("pur_sub_order_id", _purSubOrderId)
	return nil
}

// GetPurSubOrderId PurSubOrderId Getter
func (r TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest) GetPurSubOrderId() int64 {
	return r._purSubOrderId
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}
