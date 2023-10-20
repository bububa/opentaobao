package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocaipiaoshopinfoinputAPIRequest 录入参加送彩票店铺信息 API请求
// taobao.caipiao.shop.info.input
//
// 录入参加送彩票店铺信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据卖家id与赠送类型（sellerid_presenttype_uk）来决定是新增数据还是修改数据。
type TaobaocaipiaoshopinfoinputAPIRequest struct {
	model.Params
	// 店铺名称
	_shopName string
	// 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
	_actStartDate string
	// 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
	_actEndDate string
	// 店铺参加的送彩票活动描述
	_shopDesc string
	// 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空
	_presentType int64
	// 店铺类目编号，不可为空
	_shopType int64
}

// NewTaobaocaipiaoshopinfoinputRequest 初始化TaobaocaipiaoshopinfoinputAPIRequest对象
func NewTaobaocaipiaoshopinfoinputRequest() *TaobaocaipiaoshopinfoinputAPIRequest {
	return &TaobaocaipiaoshopinfoinputAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocaipiaoshopinfoinputAPIRequest) GetApiMethodName() string {
	return "taobao.caipiao.shop.info.input"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocaipiaoshopinfoinputAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocaipiaoshopinfoinputAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopName is ShopName Setter
// 店铺名称
func (r *TaobaocaipiaoshopinfoinputAPIRequest) SetShopName(_shopName string) error {
	r._shopName = _shopName
	r.Set("shop_name", _shopName)
	return nil
}

// GetShopName ShopName Getter
func (r TaobaocaipiaoshopinfoinputAPIRequest) GetShopName() string {
	return r._shopName
}

// SetActStartDate is ActStartDate Setter
// 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
func (r *TaobaocaipiaoshopinfoinputAPIRequest) SetActStartDate(_actStartDate string) error {
	r._actStartDate = _actStartDate
	r.Set("act_start_date", _actStartDate)
	return nil
}

// GetActStartDate ActStartDate Getter
func (r TaobaocaipiaoshopinfoinputAPIRequest) GetActStartDate() string {
	return r._actStartDate
}

// SetActEndDate is ActEndDate Setter
// 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
func (r *TaobaocaipiaoshopinfoinputAPIRequest) SetActEndDate(_actEndDate string) error {
	r._actEndDate = _actEndDate
	r.Set("act_end_date", _actEndDate)
	return nil
}

// GetActEndDate ActEndDate Getter
func (r TaobaocaipiaoshopinfoinputAPIRequest) GetActEndDate() string {
	return r._actEndDate
}

// SetShopDesc is ShopDesc Setter
// 店铺参加的送彩票活动描述
func (r *TaobaocaipiaoshopinfoinputAPIRequest) SetShopDesc(_shopDesc string) error {
	r._shopDesc = _shopDesc
	r.Set("shop_desc", _shopDesc)
	return nil
}

// GetShopDesc ShopDesc Getter
func (r TaobaocaipiaoshopinfoinputAPIRequest) GetShopDesc() string {
	return r._shopDesc
}

// SetPresentType is PresentType Setter
// 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空
func (r *TaobaocaipiaoshopinfoinputAPIRequest) SetPresentType(_presentType int64) error {
	r._presentType = _presentType
	r.Set("present_type", _presentType)
	return nil
}

// GetPresentType PresentType Getter
func (r TaobaocaipiaoshopinfoinputAPIRequest) GetPresentType() int64 {
	return r._presentType
}

// SetShopType is ShopType Setter
// 店铺类目编号，不可为空
func (r *TaobaocaipiaoshopinfoinputAPIRequest) SetShopType(_shopType int64) error {
	r._shopType = _shopType
	r.Set("shop_type", _shopType)
	return nil
}

// GetShopType ShopType Getter
func (r TaobaocaipiaoshopinfoinputAPIRequest) GetShopType() int64 {
	return r._shopType
}
