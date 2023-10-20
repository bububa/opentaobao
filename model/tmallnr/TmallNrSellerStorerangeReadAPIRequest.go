package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrsellerstorerangereadAPIRequest 门店服务范围读取 API请求
// tmall.nr.seller.storerange.read
//
// 读取卖家所属门店的服务范围
type TmallnrsellerstorerangereadAPIRequest struct {
	model.Params
	// 门店id
	_storeIds []int64
	// 业务身份，此api非必填
	_bizIdentity string
	// 给ISV用sellerid，ISV可能对接其他seller，有可能和登录不是同一个sellerid
	_sellerId int64
}

// NewTmallnrsellerstorerangereadRequest 初始化TmallnrsellerstorerangereadAPIRequest对象
func NewTmallnrsellerstorerangereadRequest() *TmallnrsellerstorerangereadAPIRequest {
	return &TmallnrsellerstorerangereadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrsellerstorerangereadAPIRequest) GetApiMethodName() string {
	return "tmall.nr.seller.storerange.read"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrsellerstorerangereadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrsellerstorerangereadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreIds is StoreIds Setter
// 门店id
func (r *TmallnrsellerstorerangereadAPIRequest) SetStoreIds(_storeIds []int64) error {
	r._storeIds = _storeIds
	r.Set("store_ids", _storeIds)
	return nil
}

// GetStoreIds StoreIds Getter
func (r TmallnrsellerstorerangereadAPIRequest) GetStoreIds() []int64 {
	return r._storeIds
}

// SetBizIdentity is BizIdentity Setter
// 业务身份，此api非必填
func (r *TmallnrsellerstorerangereadAPIRequest) SetBizIdentity(_bizIdentity string) error {
	r._bizIdentity = _bizIdentity
	r.Set("biz_identity", _bizIdentity)
	return nil
}

// GetBizIdentity BizIdentity Getter
func (r TmallnrsellerstorerangereadAPIRequest) GetBizIdentity() string {
	return r._bizIdentity
}

// SetSellerId is SellerId Setter
// 给ISV用sellerid，ISV可能对接其他seller，有可能和登录不是同一个sellerid
func (r *TmallnrsellerstorerangereadAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TmallnrsellerstorerangereadAPIRequest) GetSellerId() int64 {
	return r._sellerId
}
