package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrSellerStorerangeReadAPIRequest 门店服务范围读取 API请求
// tmall.nr.seller.storerange.read
//
// 读取卖家所属门店的服务范围
type TmallNrSellerStorerangeReadAPIRequest struct {
	model.Params
	// 门店id
	_storeIds []int64
	// 业务身份，此api非必填
	_bizIdentity string
	// 给ISV用sellerid，ISV可能对接其他seller，有可能和登录不是同一个sellerid
	_sellerId int64
}

// NewTmallNrSellerStorerangeReadRequest 初始化TmallNrSellerStorerangeReadAPIRequest对象
func NewTmallNrSellerStorerangeReadRequest() *TmallNrSellerStorerangeReadAPIRequest {
	return &TmallNrSellerStorerangeReadAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrSellerStorerangeReadAPIRequest) Reset() {
	r._storeIds = r._storeIds[:0]
	r._bizIdentity = ""
	r._sellerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrSellerStorerangeReadAPIRequest) GetApiMethodName() string {
	return "tmall.nr.seller.storerange.read"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrSellerStorerangeReadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrSellerStorerangeReadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreIds is StoreIds Setter
// 门店id
func (r *TmallNrSellerStorerangeReadAPIRequest) SetStoreIds(_storeIds []int64) error {
	r._storeIds = _storeIds
	r.Set("store_ids", _storeIds)
	return nil
}

// GetStoreIds StoreIds Getter
func (r TmallNrSellerStorerangeReadAPIRequest) GetStoreIds() []int64 {
	return r._storeIds
}

// SetBizIdentity is BizIdentity Setter
// 业务身份，此api非必填
func (r *TmallNrSellerStorerangeReadAPIRequest) SetBizIdentity(_bizIdentity string) error {
	r._bizIdentity = _bizIdentity
	r.Set("biz_identity", _bizIdentity)
	return nil
}

// GetBizIdentity BizIdentity Getter
func (r TmallNrSellerStorerangeReadAPIRequest) GetBizIdentity() string {
	return r._bizIdentity
}

// SetSellerId is SellerId Setter
// 给ISV用sellerid，ISV可能对接其他seller，有可能和登录不是同一个sellerid
func (r *TmallNrSellerStorerangeReadAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TmallNrSellerStorerangeReadAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

var poolTmallNrSellerStorerangeReadAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrSellerStorerangeReadRequest()
	},
}

// GetTmallNrSellerStorerangeReadRequest 从 sync.Pool 获取 TmallNrSellerStorerangeReadAPIRequest
func GetTmallNrSellerStorerangeReadAPIRequest() *TmallNrSellerStorerangeReadAPIRequest {
	return poolTmallNrSellerStorerangeReadAPIRequest.Get().(*TmallNrSellerStorerangeReadAPIRequest)
}

// ReleaseTmallNrSellerStorerangeReadAPIRequest 将 TmallNrSellerStorerangeReadAPIRequest 放入 sync.Pool
func ReleaseTmallNrSellerStorerangeReadAPIRequest(v *TmallNrSellerStorerangeReadAPIRequest) {
	v.Reset()
	poolTmallNrSellerStorerangeReadAPIRequest.Put(v)
}
