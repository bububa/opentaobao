package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrSellerStorerangeReadAPIRequest
门店服务范围读取 API请求
tmall.nr.seller.storerange.read

读取卖家所属门店的服务范围 */
type TmallNrSellerStorerangeReadAPIRequest struct {
	model.Params
	// 给ISV用sellerid，ISV可能对接其他seller，有可能和登录不是同一个sellerid
	_sellerId int64
	// 业务身份，此api非必填
	_bizIdentity string
	// 门店id
	_storeIds []int64
}

// NewTmallNrSellerStorerangeReadRequest 初始化TmallNrSellerStorerangeReadAPIRequest对象
func NewTmallNrSellerStorerangeReadRequest() *TmallNrSellerStorerangeReadAPIRequest {
	return &TmallNrSellerStorerangeReadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrSellerStorerangeReadAPIRequest) GetApiMethodName() string {
	return "tmall.nr.seller.storerange.read"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrSellerStorerangeReadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SellerId Setter
// 给ISV用sellerid，ISV可能对接其他seller，有可能和登录不是同一个sellerid
func (r *TmallNrSellerStorerangeReadAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// Get SellerId Getter
func (r TmallNrSellerStorerangeReadAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// Set is BizIdentity Setter
// 业务身份，此api非必填
func (r *TmallNrSellerStorerangeReadAPIRequest) SetBizIdentity(_bizIdentity string) error {
	r._bizIdentity = _bizIdentity
	r.Set("biz_identity", _bizIdentity)
	return nil
}

// Get BizIdentity Getter
func (r TmallNrSellerStorerangeReadAPIRequest) GetBizIdentity() string {
	return r._bizIdentity
}

// Set is StoreIds Setter
// 门店id
func (r *TmallNrSellerStorerangeReadAPIRequest) SetStoreIds(_storeIds []int64) error {
	r._storeIds = _storeIds
	r.Set("store_ids", _storeIds)
	return nil
}

// Get StoreIds Getter
func (r TmallNrSellerStorerangeReadAPIRequest) GetStoreIds() []int64 {
	return r._storeIds
}
