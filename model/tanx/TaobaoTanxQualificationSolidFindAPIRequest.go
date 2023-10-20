package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotanxqualificationsolidfindAPIRequest 客户固态共享资质 API请求
// taobao.tanx.qualification.solid.find
//
// 接口会返回该广告主下的所有审核通过并且可被共享的资质，这些资质在过期之前可以不需要再次上传。
type TaobaotanxqualificationsolidfindAPIRequest struct {
	model.Params
	// 资质元素id列表
	_elementIds []int64
	// dsp客户验证token
	_token string
	// 广告主id
	_advertiserId int64
	// dsp用户id
	_memberId int64
	// 1970年到现在的秒
	_signTime int64
	// 查询起始页
	_page int64
	// 分页大小
	_pageSize int64
}

// NewTaobaotanxqualificationsolidfindRequest 初始化TaobaotanxqualificationsolidfindAPIRequest对象
func NewTaobaotanxqualificationsolidfindRequest() *TaobaotanxqualificationsolidfindAPIRequest {
	return &TaobaotanxqualificationsolidfindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotanxqualificationsolidfindAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.qualification.solid.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotanxqualificationsolidfindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotanxqualificationsolidfindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetElementIds is ElementIds Setter
// 资质元素id列表
func (r *TaobaotanxqualificationsolidfindAPIRequest) SetElementIds(_elementIds []int64) error {
	r._elementIds = _elementIds
	r.Set("element_ids", _elementIds)
	return nil
}

// GetElementIds ElementIds Getter
func (r TaobaotanxqualificationsolidfindAPIRequest) GetElementIds() []int64 {
	return r._elementIds
}

// SetToken is Token Setter
// dsp客户验证token
func (r *TaobaotanxqualificationsolidfindAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaotanxqualificationsolidfindAPIRequest) GetToken() string {
	return r._token
}

// SetAdvertiserId is AdvertiserId Setter
// 广告主id
func (r *TaobaotanxqualificationsolidfindAPIRequest) SetAdvertiserId(_advertiserId int64) error {
	r._advertiserId = _advertiserId
	r.Set("advertiser_id", _advertiserId)
	return nil
}

// GetAdvertiserId AdvertiserId Getter
func (r TaobaotanxqualificationsolidfindAPIRequest) GetAdvertiserId() int64 {
	return r._advertiserId
}

// SetMemberId is MemberId Setter
// dsp用户id
func (r *TaobaotanxqualificationsolidfindAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaotanxqualificationsolidfindAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetSignTime is SignTime Setter
// 1970年到现在的秒
func (r *TaobaotanxqualificationsolidfindAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaotanxqualificationsolidfindAPIRequest) GetSignTime() int64 {
	return r._signTime
}

// SetPage is Page Setter
// 查询起始页
func (r *TaobaotanxqualificationsolidfindAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r TaobaotanxqualificationsolidfindAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TaobaotanxqualificationsolidfindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaotanxqualificationsolidfindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
