package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScOptimusPromotionAPIRequest 淘宝客-服务商-权益物料精选 API请求
// taobao.tbk.sc.optimus.promotion
//
// 服务商使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
type TaobaoTbkScOptimusPromotionAPIRequest struct {
	model.Params
	// 页大小，每次请求限制10以内
	_pageSize int64
	// 第几页，默认：1
	_pageNum int64
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneId int64
	// 官方提供的权益物料Id。有价券-37104、大额店铺券-37116、天猫店铺券-62191、券券补-61809 更多权益物料id敬请期待！
	_promotionId int64
	// mm_xxx_xxx_xxx的第2段数字
	_siteId int64
}

// NewTaobaoTbkScOptimusPromotionRequest 初始化TaobaoTbkScOptimusPromotionAPIRequest对象
func NewTaobaoTbkScOptimusPromotionRequest() *TaobaoTbkScOptimusPromotionAPIRequest {
	return &TaobaoTbkScOptimusPromotionAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkScOptimusPromotionAPIRequest) Reset() {
	r._pageSize = 0
	r._pageNum = 0
	r._adzoneId = 0
	r._promotionId = 0
	r._siteId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScOptimusPromotionAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.optimus.promotion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScOptimusPromotionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScOptimusPromotionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageSize is PageSize Setter
// 页大小，每次请求限制10以内
func (r *TaobaoTbkScOptimusPromotionAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTbkScOptimusPromotionAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 第几页，默认：1
func (r *TaobaoTbkScOptimusPromotionAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoTbkScOptimusPromotionAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetAdzoneId is AdzoneId Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaoTbkScOptimusPromotionAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkScOptimusPromotionAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetPromotionId is PromotionId Setter
// 官方提供的权益物料Id。有价券-37104、大额店铺券-37116、天猫店铺券-62191、券券补-61809 更多权益物料id敬请期待！
func (r *TaobaoTbkScOptimusPromotionAPIRequest) SetPromotionId(_promotionId int64) error {
	r._promotionId = _promotionId
	r.Set("promotion_id", _promotionId)
	return nil
}

// GetPromotionId PromotionId Getter
func (r TaobaoTbkScOptimusPromotionAPIRequest) GetPromotionId() int64 {
	return r._promotionId
}

// SetSiteId is SiteId Setter
// mm_xxx_xxx_xxx的第2段数字
func (r *TaobaoTbkScOptimusPromotionAPIRequest) SetSiteId(_siteId int64) error {
	r._siteId = _siteId
	r.Set("site_id", _siteId)
	return nil
}

// GetSiteId SiteId Getter
func (r TaobaoTbkScOptimusPromotionAPIRequest) GetSiteId() int64 {
	return r._siteId
}

var poolTaobaoTbkScOptimusPromotionAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkScOptimusPromotionRequest()
	},
}

// GetTaobaoTbkScOptimusPromotionRequest 从 sync.Pool 获取 TaobaoTbkScOptimusPromotionAPIRequest
func GetTaobaoTbkScOptimusPromotionAPIRequest() *TaobaoTbkScOptimusPromotionAPIRequest {
	return poolTaobaoTbkScOptimusPromotionAPIRequest.Get().(*TaobaoTbkScOptimusPromotionAPIRequest)
}

// ReleaseTaobaoTbkScOptimusPromotionAPIRequest 将 TaobaoTbkScOptimusPromotionAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkScOptimusPromotionAPIRequest(v *TaobaoTbkScOptimusPromotionAPIRequest) {
	v.Reset()
	poolTaobaoTbkScOptimusPromotionAPIRequest.Put(v)
}
