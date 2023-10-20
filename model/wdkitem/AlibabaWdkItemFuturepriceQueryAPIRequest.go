package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemfuturepricequeryAPIRequest 单个商品未来价查询接口 API请求
// alibaba.wdk.item.futureprice.query
//
// 查询单个商品未来价，融合了未来基础售价+未来促销价
type AlibabawdkitemfuturepricequeryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 渠道
	_orderChannelCode string
	// 开始时间
	_startTime string
	// 结束时间，结束时间-开始时间不能超过48小时
	_endTime string
	// 渠道店id
	_shopId int64
}

// NewAlibabawdkitemfuturepricequeryRequest 初始化AlibabawdkitemfuturepricequeryAPIRequest对象
func NewAlibabawdkitemfuturepricequeryRequest() *AlibabawdkitemfuturepricequeryAPIRequest {
	return &AlibabawdkitemfuturepricequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemfuturepricequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.futureprice.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemfuturepricequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemfuturepricequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabawdkitemfuturepricequeryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabawdkitemfuturepricequeryAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetOrderChannelCode is OrderChannelCode Setter
// 渠道
func (r *AlibabawdkitemfuturepricequeryAPIRequest) SetOrderChannelCode(_orderChannelCode string) error {
	r._orderChannelCode = _orderChannelCode
	r.Set("order_channel_code", _orderChannelCode)
	return nil
}

// GetOrderChannelCode OrderChannelCode Getter
func (r AlibabawdkitemfuturepricequeryAPIRequest) GetOrderChannelCode() string {
	return r._orderChannelCode
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *AlibabawdkitemfuturepricequeryAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AlibabawdkitemfuturepricequeryAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间，结束时间-开始时间不能超过48小时
func (r *AlibabawdkitemfuturepricequeryAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabawdkitemfuturepricequeryAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetShopId is ShopId Setter
// 渠道店id
func (r *AlibabawdkitemfuturepricequeryAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabawdkitemfuturepricequeryAPIRequest) GetShopId() int64 {
	return r._shopId
}
