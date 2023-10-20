package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotioncouponsellersearchAPIRequest 查询绑定卖家优惠券相关信息 API请求
// taobao.promotion.coupon.seller.search
//
// 查询绑定卖家相关优惠券信息  如isv  百川 等外部业务方
type TaobaopromotioncouponsellersearchAPIRequest struct {
	model.Params
	// 券id集合
	_spreadIds []string
	// 卖家昵称
	_sellerNick string
	// 当前第几页  从第一页开始
	_currentPage int64
	// 每页数据 最大20左右
	_pageSize int64
}

// NewTaobaopromotioncouponsellersearchRequest 初始化TaobaopromotioncouponsellersearchAPIRequest对象
func NewTaobaopromotioncouponsellersearchRequest() *TaobaopromotioncouponsellersearchAPIRequest {
	return &TaobaopromotioncouponsellersearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotioncouponsellersearchAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.coupon.seller.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotioncouponsellersearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotioncouponsellersearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSpreadIds is SpreadIds Setter
// 券id集合
func (r *TaobaopromotioncouponsellersearchAPIRequest) SetSpreadIds(_spreadIds []string) error {
	r._spreadIds = _spreadIds
	r.Set("spread_ids", _spreadIds)
	return nil
}

// GetSpreadIds SpreadIds Getter
func (r TaobaopromotioncouponsellersearchAPIRequest) GetSpreadIds() []string {
	return r._spreadIds
}

// SetSellerNick is SellerNick Setter
// 卖家昵称
func (r *TaobaopromotioncouponsellersearchAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaopromotioncouponsellersearchAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// SetCurrentPage is CurrentPage Setter
// 当前第几页  从第一页开始
func (r *TaobaopromotioncouponsellersearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaopromotioncouponsellersearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页数据 最大20左右
func (r *TaobaopromotioncouponsellersearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaopromotioncouponsellersearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
