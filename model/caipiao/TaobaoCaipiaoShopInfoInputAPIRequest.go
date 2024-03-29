package caipiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCaipiaoShopInfoInputAPIRequest 录入参加送彩票店铺信息 API请求
// taobao.caipiao.shop.info.input
//
// 录入参加送彩票店铺信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据卖家id与赠送类型（sellerid_presenttype_uk）来决定是新增数据还是修改数据。
type TaobaoCaipiaoShopInfoInputAPIRequest struct {
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

// NewTaobaoCaipiaoShopInfoInputRequest 初始化TaobaoCaipiaoShopInfoInputAPIRequest对象
func NewTaobaoCaipiaoShopInfoInputRequest() *TaobaoCaipiaoShopInfoInputAPIRequest {
	return &TaobaoCaipiaoShopInfoInputAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCaipiaoShopInfoInputAPIRequest) Reset() {
	r._shopName = ""
	r._actStartDate = ""
	r._actEndDate = ""
	r._shopDesc = ""
	r._presentType = 0
	r._shopType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoShopInfoInputAPIRequest) GetApiMethodName() string {
	return "taobao.caipiao.shop.info.input"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoShopInfoInputAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCaipiaoShopInfoInputAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopName is ShopName Setter
// 店铺名称
func (r *TaobaoCaipiaoShopInfoInputAPIRequest) SetShopName(_shopName string) error {
	r._shopName = _shopName
	r.Set("shop_name", _shopName)
	return nil
}

// GetShopName ShopName Getter
func (r TaobaoCaipiaoShopInfoInputAPIRequest) GetShopName() string {
	return r._shopName
}

// SetActStartDate is ActStartDate Setter
// 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
func (r *TaobaoCaipiaoShopInfoInputAPIRequest) SetActStartDate(_actStartDate string) error {
	r._actStartDate = _actStartDate
	r.Set("act_start_date", _actStartDate)
	return nil
}

// GetActStartDate ActStartDate Getter
func (r TaobaoCaipiaoShopInfoInputAPIRequest) GetActStartDate() string {
	return r._actStartDate
}

// SetActEndDate is ActEndDate Setter
// 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
func (r *TaobaoCaipiaoShopInfoInputAPIRequest) SetActEndDate(_actEndDate string) error {
	r._actEndDate = _actEndDate
	r.Set("act_end_date", _actEndDate)
	return nil
}

// GetActEndDate ActEndDate Getter
func (r TaobaoCaipiaoShopInfoInputAPIRequest) GetActEndDate() string {
	return r._actEndDate
}

// SetShopDesc is ShopDesc Setter
// 店铺参加的送彩票活动描述
func (r *TaobaoCaipiaoShopInfoInputAPIRequest) SetShopDesc(_shopDesc string) error {
	r._shopDesc = _shopDesc
	r.Set("shop_desc", _shopDesc)
	return nil
}

// GetShopDesc ShopDesc Getter
func (r TaobaoCaipiaoShopInfoInputAPIRequest) GetShopDesc() string {
	return r._shopDesc
}

// SetPresentType is PresentType Setter
// 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空
func (r *TaobaoCaipiaoShopInfoInputAPIRequest) SetPresentType(_presentType int64) error {
	r._presentType = _presentType
	r.Set("present_type", _presentType)
	return nil
}

// GetPresentType PresentType Getter
func (r TaobaoCaipiaoShopInfoInputAPIRequest) GetPresentType() int64 {
	return r._presentType
}

// SetShopType is ShopType Setter
// 店铺类目编号，不可为空
func (r *TaobaoCaipiaoShopInfoInputAPIRequest) SetShopType(_shopType int64) error {
	r._shopType = _shopType
	r.Set("shop_type", _shopType)
	return nil
}

// GetShopType ShopType Getter
func (r TaobaoCaipiaoShopInfoInputAPIRequest) GetShopType() int64 {
	return r._shopType
}

var poolTaobaoCaipiaoShopInfoInputAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCaipiaoShopInfoInputRequest()
	},
}

// GetTaobaoCaipiaoShopInfoInputRequest 从 sync.Pool 获取 TaobaoCaipiaoShopInfoInputAPIRequest
func GetTaobaoCaipiaoShopInfoInputAPIRequest() *TaobaoCaipiaoShopInfoInputAPIRequest {
	return poolTaobaoCaipiaoShopInfoInputAPIRequest.Get().(*TaobaoCaipiaoShopInfoInputAPIRequest)
}

// ReleaseTaobaoCaipiaoShopInfoInputAPIRequest 将 TaobaoCaipiaoShopInfoInputAPIRequest 放入 sync.Pool
func ReleaseTaobaoCaipiaoShopInfoInputAPIRequest(v *TaobaoCaipiaoShopInfoInputAPIRequest) {
	v.Reset()
	poolTaobaoCaipiaoShopInfoInputAPIRequest.Put(v)
}
