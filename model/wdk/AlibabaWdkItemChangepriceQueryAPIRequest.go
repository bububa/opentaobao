package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemChangepriceQueryAPIRequest 按照价格变更时间段，查询会变更价格的单据的商品 API请求
// alibaba.wdk.item.changeprice.query
//
// *
//      * 按照价格变更时间段，查询会变更价格的单据的商品
//      * 传入QueryPriceChangeTypeEnum.BASE_PRICE,返回变价时间在startTime-endTime之间的所有单据
//      * 传入QueryPriceChangeTypeEnum.SKU_PROMOTION_START,
//      * 返回活动开始时间在 startTime<=活动开始时间<endTime 之间的所有单品促销单据
//      * 传入QueryPriceChangeTypeEnum.SKU_PROMOTION_END,
//      * 返回活动结束时间在 startTime<=活动结束时间<endTime 之间的所有单品促销单据
//
type AlibabaWdkItemChangepriceQueryAPIRequest struct {
	model.Params
	// 变价的类型  * 查询变价的单据专用
	_type string
	// 开始时间
	_startTime string
	// 结束时间，结束时间-开始时间不能超过48小时
	_endTime string
	// 渠道店id
	_shopId string
}

// NewAlibabaWdkItemChangepriceQueryRequest 初始化AlibabaWdkItemChangepriceQueryAPIRequest对象
func NewAlibabaWdkItemChangepriceQueryRequest() *AlibabaWdkItemChangepriceQueryAPIRequest {
	return &AlibabaWdkItemChangepriceQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemChangepriceQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.changeprice.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemChangepriceQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Type Setter
// 变价的类型  * 查询变价的单据专用
func (r *AlibabaWdkItemChangepriceQueryAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r AlibabaWdkItemChangepriceQueryAPIRequest) GetType() string {
	return r._type
}

// Set is StartTime Setter
// 开始时间
func (r *AlibabaWdkItemChangepriceQueryAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r AlibabaWdkItemChangepriceQueryAPIRequest) GetStartTime() string {
	return r._startTime
}

// Set is EndTime Setter
// 结束时间，结束时间-开始时间不能超过48小时
func (r *AlibabaWdkItemChangepriceQueryAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r AlibabaWdkItemChangepriceQueryAPIRequest) GetEndTime() string {
	return r._endTime
}

// Set is ShopId Setter
// 渠道店id
func (r *AlibabaWdkItemChangepriceQueryAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// Get ShopId Getter
func (r AlibabaWdkItemChangepriceQueryAPIRequest) GetShopId() string {
	return r._shopId
}
