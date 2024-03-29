package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFuwuServiceitemListAPIRequest 获取服务商品扩展信息 API请求
// tmall.fuwu.serviceitem.list
//
// 获取服务商品扩展信息
type TmallFuwuServiceitemListAPIRequest struct {
	model.Params
	// 商品id列表，有数量限制
	_itemids []string
	// 商品所属卖家账号id
	_sellerId int64
}

// NewTmallFuwuServiceitemListRequest 初始化TmallFuwuServiceitemListAPIRequest对象
func NewTmallFuwuServiceitemListRequest() *TmallFuwuServiceitemListAPIRequest {
	return &TmallFuwuServiceitemListAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallFuwuServiceitemListAPIRequest) Reset() {
	r._itemids = r._itemids[:0]
	r._sellerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallFuwuServiceitemListAPIRequest) GetApiMethodName() string {
	return "tmall.fuwu.serviceitem.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallFuwuServiceitemListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallFuwuServiceitemListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemids is Itemids Setter
// 商品id列表，有数量限制
func (r *TmallFuwuServiceitemListAPIRequest) SetItemids(_itemids []string) error {
	r._itemids = _itemids
	r.Set("itemids", _itemids)
	return nil
}

// GetItemids Itemids Getter
func (r TmallFuwuServiceitemListAPIRequest) GetItemids() []string {
	return r._itemids
}

// SetSellerId is SellerId Setter
// 商品所属卖家账号id
func (r *TmallFuwuServiceitemListAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TmallFuwuServiceitemListAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

var poolTmallFuwuServiceitemListAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallFuwuServiceitemListRequest()
	},
}

// GetTmallFuwuServiceitemListRequest 从 sync.Pool 获取 TmallFuwuServiceitemListAPIRequest
func GetTmallFuwuServiceitemListAPIRequest() *TmallFuwuServiceitemListAPIRequest {
	return poolTmallFuwuServiceitemListAPIRequest.Get().(*TmallFuwuServiceitemListAPIRequest)
}

// ReleaseTmallFuwuServiceitemListAPIRequest 将 TmallFuwuServiceitemListAPIRequest 放入 sync.Pool
func ReleaseTmallFuwuServiceitemListAPIRequest(v *TmallFuwuServiceitemListAPIRequest) {
	v.Reset()
	poolTmallFuwuServiceitemListAPIRequest.Put(v)
}
