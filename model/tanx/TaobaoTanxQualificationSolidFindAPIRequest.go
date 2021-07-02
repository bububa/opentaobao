package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxQualificationSolidFindAPIRequest 客户固态共享资质 API请求
// taobao.tanx.qualification.solid.find
//
// 接口会返回该广告主下的所有审核通过并且可被共享的资质，这些资质在过期之前可以不需要再次上传。
type TaobaoTanxQualificationSolidFindAPIRequest struct {
	model.Params
	// 广告主id
	_advertiserId int64
	// 资质元素id列表
	_elementIds []int64
	// dsp用户id
	_memberId int64
	// dsp客户验证token
	_token string
	// 1970年到现在的秒
	_signTime int64
	// 查询起始页
	_page int64
	// 分页大小
	_pageSize int64
}

// NewTaobaoTanxQualificationSolidFindRequest 初始化TaobaoTanxQualificationSolidFindAPIRequest对象
func NewTaobaoTanxQualificationSolidFindRequest() *TaobaoTanxQualificationSolidFindAPIRequest {
	return &TaobaoTanxQualificationSolidFindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTanxQualificationSolidFindAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.qualification.solid.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTanxQualificationSolidFindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAdvertiserId is AdvertiserId Setter
// 广告主id
func (r *TaobaoTanxQualificationSolidFindAPIRequest) SetAdvertiserId(_advertiserId int64) error {
	r._advertiserId = _advertiserId
	r.Set("advertiser_id", _advertiserId)
	return nil
}

// GetAdvertiserId AdvertiserId Getter
func (r TaobaoTanxQualificationSolidFindAPIRequest) GetAdvertiserId() int64 {
	return r._advertiserId
}

// SetElementIds is ElementIds Setter
// 资质元素id列表
func (r *TaobaoTanxQualificationSolidFindAPIRequest) SetElementIds(_elementIds []int64) error {
	r._elementIds = _elementIds
	r.Set("element_ids", _elementIds)
	return nil
}

// GetElementIds ElementIds Getter
func (r TaobaoTanxQualificationSolidFindAPIRequest) GetElementIds() []int64 {
	return r._elementIds
}

// SetMemberId is MemberId Setter
// dsp用户id
func (r *TaobaoTanxQualificationSolidFindAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaoTanxQualificationSolidFindAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetToken is Token Setter
// dsp客户验证token
func (r *TaobaoTanxQualificationSolidFindAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoTanxQualificationSolidFindAPIRequest) GetToken() string {
	return r._token
}

// SetSignTime is SignTime Setter
// 1970年到现在的秒
func (r *TaobaoTanxQualificationSolidFindAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaoTanxQualificationSolidFindAPIRequest) GetSignTime() int64 {
	return r._signTime
}

// SetPage is Page Setter
// 查询起始页
func (r *TaobaoTanxQualificationSolidFindAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r TaobaoTanxQualificationSolidFindAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TaobaoTanxQualificationSolidFindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTanxQualificationSolidFindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
