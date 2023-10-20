package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrSellerStorerangeSyncAPIRequest 同步商户中心服务范围 API请求
// tmall.nr.seller.storerange.sync
//
// 同步商户中心服务范围
type TmallNrSellerStorerangeSyncAPIRequest struct {
	model.Params
	// 系统自动生成
	_reqDTOList []SyncServiceRangeRequestDto
	// 业务身份标识,dss定时送；self_day 自配日达；self_hour 自配小时达
	_bizIdentity string
	// 卖家id，有可能和登录seller不是同一个id
	_sellerId int64
}

// NewTmallNrSellerStorerangeSyncRequest 初始化TmallNrSellerStorerangeSyncAPIRequest对象
func NewTmallNrSellerStorerangeSyncRequest() *TmallNrSellerStorerangeSyncAPIRequest {
	return &TmallNrSellerStorerangeSyncAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrSellerStorerangeSyncAPIRequest) Reset() {
	r._reqDTOList = r._reqDTOList[:0]
	r._bizIdentity = ""
	r._sellerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrSellerStorerangeSyncAPIRequest) GetApiMethodName() string {
	return "tmall.nr.seller.storerange.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrSellerStorerangeSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrSellerStorerangeSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReqDTOList is ReqDTOList Setter
// 系统自动生成
func (r *TmallNrSellerStorerangeSyncAPIRequest) SetReqDTOList(_reqDTOList []SyncServiceRangeRequestDto) error {
	r._reqDTOList = _reqDTOList
	r.Set("req_d_t_o_list", _reqDTOList)
	return nil
}

// GetReqDTOList ReqDTOList Getter
func (r TmallNrSellerStorerangeSyncAPIRequest) GetReqDTOList() []SyncServiceRangeRequestDto {
	return r._reqDTOList
}

// SetBizIdentity is BizIdentity Setter
// 业务身份标识,dss定时送；self_day 自配日达；self_hour 自配小时达
func (r *TmallNrSellerStorerangeSyncAPIRequest) SetBizIdentity(_bizIdentity string) error {
	r._bizIdentity = _bizIdentity
	r.Set("biz_identity", _bizIdentity)
	return nil
}

// GetBizIdentity BizIdentity Getter
func (r TmallNrSellerStorerangeSyncAPIRequest) GetBizIdentity() string {
	return r._bizIdentity
}

// SetSellerId is SellerId Setter
// 卖家id，有可能和登录seller不是同一个id
func (r *TmallNrSellerStorerangeSyncAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TmallNrSellerStorerangeSyncAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

var poolTmallNrSellerStorerangeSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrSellerStorerangeSyncRequest()
	},
}

// GetTmallNrSellerStorerangeSyncRequest 从 sync.Pool 获取 TmallNrSellerStorerangeSyncAPIRequest
func GetTmallNrSellerStorerangeSyncAPIRequest() *TmallNrSellerStorerangeSyncAPIRequest {
	return poolTmallNrSellerStorerangeSyncAPIRequest.Get().(*TmallNrSellerStorerangeSyncAPIRequest)
}

// ReleaseTmallNrSellerStorerangeSyncAPIRequest 将 TmallNrSellerStorerangeSyncAPIRequest 放入 sync.Pool
func ReleaseTmallNrSellerStorerangeSyncAPIRequest(v *TmallNrSellerStorerangeSyncAPIRequest) {
	v.Reset()
	poolTmallNrSellerStorerangeSyncAPIRequest.Put(v)
}
