package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslBizApSearchAPIRequest
AP列表查询 API请求
taobao.uscesl.biz.ap.search

查询当前门店下登记的AP列表 */
type TaobaoUsceslBizApSearchAPIRequest struct {
	model.Params
	// 商家编码
	_bizBrandKey string
	// 每页显示数
	_limit int64
	// 是否激活
	_isActivate bool
	// 价签条码
	_mac string
	// 页码
	_currentPage int64
	// 门店ID
	_storeId int64
}

// NewTaobaoUsceslBizApSearchRequest 初始化TaobaoUsceslBizApSearchAPIRequest对象
func NewTaobaoUsceslBizApSearchRequest() *TaobaoUsceslBizApSearchAPIRequest {
	return &TaobaoUsceslBizApSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizApSearchAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.ap.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizApSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizBrandKey Setter
// 商家编码
func (r *TaobaoUsceslBizApSearchAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// Get BizBrandKey Getter
func (r TaobaoUsceslBizApSearchAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// Set is Limit Setter
// 每页显示数
func (r *TaobaoUsceslBizApSearchAPIRequest) SetLimit(_limit int64) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// Get Limit Getter
func (r TaobaoUsceslBizApSearchAPIRequest) GetLimit() int64 {
	return r._limit
}

// Set is IsActivate Setter
// 是否激活
func (r *TaobaoUsceslBizApSearchAPIRequest) SetIsActivate(_isActivate bool) error {
	r._isActivate = _isActivate
	r.Set("is_activate", _isActivate)
	return nil
}

// Get IsActivate Getter
func (r TaobaoUsceslBizApSearchAPIRequest) GetIsActivate() bool {
	return r._isActivate
}

// Set is Mac Setter
// 价签条码
func (r *TaobaoUsceslBizApSearchAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// Get Mac Getter
func (r TaobaoUsceslBizApSearchAPIRequest) GetMac() string {
	return r._mac
}

// Set is CurrentPage Setter
// 页码
func (r *TaobaoUsceslBizApSearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// Get CurrentPage Getter
func (r TaobaoUsceslBizApSearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// Set is StoreId Setter
// 门店ID
func (r *TaobaoUsceslBizApSearchAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoUsceslBizApSearchAPIRequest) GetStoreId() int64 {
	return r._storeId
}
