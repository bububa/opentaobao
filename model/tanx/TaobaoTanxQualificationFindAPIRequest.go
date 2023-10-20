package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotanxqualificationfindAPIRequest 资质查询接口 API请求
// taobao.tanx.qualification.find
//
// 资质查询接口
type TaobaotanxqualificationfindAPIRequest struct {
	model.Params
	// dsp客户在tanx签名的token
	_token string
	// dsp客户在tanx的memberId
	_memberId int64
	// 当前时间,从1970年算的毫秒
	_signTime int64
	// 分页的第几页，从1开始
	_page int64
	// 分页大小，限制200
	_pageSize int64
	// 广告主资质查询dto
	_query *QualificationQuery
}

// NewTaobaotanxqualificationfindRequest 初始化TaobaotanxqualificationfindAPIRequest对象
func NewTaobaotanxqualificationfindRequest() *TaobaotanxqualificationfindAPIRequest {
	return &TaobaotanxqualificationfindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotanxqualificationfindAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.qualification.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotanxqualificationfindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotanxqualificationfindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// dsp客户在tanx签名的token
func (r *TaobaotanxqualificationfindAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaotanxqualificationfindAPIRequest) GetToken() string {
	return r._token
}

// SetMemberId is MemberId Setter
// dsp客户在tanx的memberId
func (r *TaobaotanxqualificationfindAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaotanxqualificationfindAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetSignTime is SignTime Setter
// 当前时间,从1970年算的毫秒
func (r *TaobaotanxqualificationfindAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaotanxqualificationfindAPIRequest) GetSignTime() int64 {
	return r._signTime
}

// SetPage is Page Setter
// 分页的第几页，从1开始
func (r *TaobaotanxqualificationfindAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r TaobaotanxqualificationfindAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 分页大小，限制200
func (r *TaobaotanxqualificationfindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaotanxqualificationfindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetQuery is Query Setter
// 广告主资质查询dto
func (r *TaobaotanxqualificationfindAPIRequest) SetQuery(_query *QualificationQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TaobaotanxqualificationfindAPIRequest) GetQuery() *QualificationQuery {
	return r._query
}
