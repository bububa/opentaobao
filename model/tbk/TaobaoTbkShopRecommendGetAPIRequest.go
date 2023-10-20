package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkShopRecommendGetAPIRequest 淘宝客-公用-店铺关联推荐 API请求
// taobao.tbk.shop.recommend.get
//
// 入参卖家id，可推荐与此店铺相关联的相关店铺。
type TaobaoTbkShopRecommendGetAPIRequest struct {
	model.Params
	// 需返回的字段列表
	_fields string
	// 返回数量，默认20，最大值40
	_count int64
	// 链接形式：1：PC，2：无线，默认：１
	_platform int64
	// 卖家Id
	_userId int64
}

// NewTaobaoTbkShopRecommendGetRequest 初始化TaobaoTbkShopRecommendGetAPIRequest对象
func NewTaobaoTbkShopRecommendGetRequest() *TaobaoTbkShopRecommendGetAPIRequest {
	return &TaobaoTbkShopRecommendGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkShopRecommendGetAPIRequest) Reset() {
	r._fields = ""
	r._count = 0
	r._platform = 0
	r._userId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkShopRecommendGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.shop.recommend.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkShopRecommendGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkShopRecommendGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表
func (r *TaobaoTbkShopRecommendGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTbkShopRecommendGetAPIRequest) GetFields() string {
	return r._fields
}

// SetCount is Count Setter
// 返回数量，默认20，最大值40
func (r *TaobaoTbkShopRecommendGetAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// GetCount Count Getter
func (r TaobaoTbkShopRecommendGetAPIRequest) GetCount() int64 {
	return r._count
}

// SetPlatform is Platform Setter
// 链接形式：1：PC，2：无线，默认：１
func (r *TaobaoTbkShopRecommendGetAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaoTbkShopRecommendGetAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetUserId is UserId Setter
// 卖家Id
func (r *TaobaoTbkShopRecommendGetAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoTbkShopRecommendGetAPIRequest) GetUserId() int64 {
	return r._userId
}

var poolTaobaoTbkShopRecommendGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkShopRecommendGetRequest()
	},
}

// GetTaobaoTbkShopRecommendGetRequest 从 sync.Pool 获取 TaobaoTbkShopRecommendGetAPIRequest
func GetTaobaoTbkShopRecommendGetAPIRequest() *TaobaoTbkShopRecommendGetAPIRequest {
	return poolTaobaoTbkShopRecommendGetAPIRequest.Get().(*TaobaoTbkShopRecommendGetAPIRequest)
}

// ReleaseTaobaoTbkShopRecommendGetAPIRequest 将 TaobaoTbkShopRecommendGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkShopRecommendGetAPIRequest(v *TaobaoTbkShopRecommendGetAPIRequest) {
	v.Reset()
	poolTaobaoTbkShopRecommendGetAPIRequest.Put(v)
}
