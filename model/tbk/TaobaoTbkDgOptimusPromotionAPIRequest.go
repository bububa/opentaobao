package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgOptimusPromotionAPIRequest 淘宝客-推广者-权益物料精选 API请求
// taobao.tbk.dg.optimus.promotion
//
// 推广者使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
type TaobaoTbkDgOptimusPromotionAPIRequest struct {
	model.Params
	// 页大小，一次请求请限制在10以内
	_pageSize int64
	// 第几页，默认：1
	_pageNum int64
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneId int64
	// 官方提供的权益物料Id。有价券-37104、大额店铺券-37116、天猫店铺券-62191、券券补-61809 更多权益物料id敬请期待！
	_promotionId int64
}

// NewTaobaoTbkDgOptimusPromotionRequest 初始化TaobaoTbkDgOptimusPromotionAPIRequest对象
func NewTaobaoTbkDgOptimusPromotionRequest() *TaobaoTbkDgOptimusPromotionAPIRequest {
	return &TaobaoTbkDgOptimusPromotionAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkDgOptimusPromotionAPIRequest) Reset() {
	r._pageSize = 0
	r._pageNum = 0
	r._adzoneId = 0
	r._promotionId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgOptimusPromotionAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.optimus.promotion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgOptimusPromotionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkDgOptimusPromotionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageSize is PageSize Setter
// 页大小，一次请求请限制在10以内
func (r *TaobaoTbkDgOptimusPromotionAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTbkDgOptimusPromotionAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 第几页，默认：1
func (r *TaobaoTbkDgOptimusPromotionAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoTbkDgOptimusPromotionAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetAdzoneId is AdzoneId Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaoTbkDgOptimusPromotionAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkDgOptimusPromotionAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetPromotionId is PromotionId Setter
// 官方提供的权益物料Id。有价券-37104、大额店铺券-37116、天猫店铺券-62191、券券补-61809 更多权益物料id敬请期待！
func (r *TaobaoTbkDgOptimusPromotionAPIRequest) SetPromotionId(_promotionId int64) error {
	r._promotionId = _promotionId
	r.Set("promotion_id", _promotionId)
	return nil
}

// GetPromotionId PromotionId Getter
func (r TaobaoTbkDgOptimusPromotionAPIRequest) GetPromotionId() int64 {
	return r._promotionId
}

var poolTaobaoTbkDgOptimusPromotionAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkDgOptimusPromotionRequest()
	},
}

// GetTaobaoTbkDgOptimusPromotionRequest 从 sync.Pool 获取 TaobaoTbkDgOptimusPromotionAPIRequest
func GetTaobaoTbkDgOptimusPromotionAPIRequest() *TaobaoTbkDgOptimusPromotionAPIRequest {
	return poolTaobaoTbkDgOptimusPromotionAPIRequest.Get().(*TaobaoTbkDgOptimusPromotionAPIRequest)
}

// ReleaseTaobaoTbkDgOptimusPromotionAPIRequest 将 TaobaoTbkDgOptimusPromotionAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkDgOptimusPromotionAPIRequest(v *TaobaoTbkDgOptimusPromotionAPIRequest) {
	v.Reset()
	poolTaobaoTbkDgOptimusPromotionAPIRequest.Put(v)
}
