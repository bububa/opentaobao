package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomevrsyncAPIRequest VR关系数据同步 API请求
// alibaba.alihouse.newhome.vr.sync
//
// 对接易居VR关系数据迁移
type AlibabaalihousenewhomevrsyncAPIRequest struct {
	model.Params
	// VR提取码
	_extractedCode string
	// 户型ID
	_layoutInfoId string
	// 是否生效
	_isValid string
	// 封面图
	_coverImage string
	// vr展示链接
	_vrUrl string
	// 营销户型图
	_markingLayoutImg string
	// 是否是单层 1 是 0 否
	_isSingleLayout int64
}

// NewAlibabaalihousenewhomevrsyncRequest 初始化AlibabaalihousenewhomevrsyncAPIRequest对象
func NewAlibabaalihousenewhomevrsyncRequest() *AlibabaalihousenewhomevrsyncAPIRequest {
	return &AlibabaalihousenewhomevrsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomevrsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.vr.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomevrsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomevrsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtractedCode is ExtractedCode Setter
// VR提取码
func (r *AlibabaalihousenewhomevrsyncAPIRequest) SetExtractedCode(_extractedCode string) error {
	r._extractedCode = _extractedCode
	r.Set("extracted_code", _extractedCode)
	return nil
}

// GetExtractedCode ExtractedCode Getter
func (r AlibabaalihousenewhomevrsyncAPIRequest) GetExtractedCode() string {
	return r._extractedCode
}

// SetLayoutInfoId is LayoutInfoId Setter
// 户型ID
func (r *AlibabaalihousenewhomevrsyncAPIRequest) SetLayoutInfoId(_layoutInfoId string) error {
	r._layoutInfoId = _layoutInfoId
	r.Set("layout_info_id", _layoutInfoId)
	return nil
}

// GetLayoutInfoId LayoutInfoId Getter
func (r AlibabaalihousenewhomevrsyncAPIRequest) GetLayoutInfoId() string {
	return r._layoutInfoId
}

// SetIsValid is IsValid Setter
// 是否生效
func (r *AlibabaalihousenewhomevrsyncAPIRequest) SetIsValid(_isValid string) error {
	r._isValid = _isValid
	r.Set("is_valid", _isValid)
	return nil
}

// GetIsValid IsValid Getter
func (r AlibabaalihousenewhomevrsyncAPIRequest) GetIsValid() string {
	return r._isValid
}

// SetCoverImage is CoverImage Setter
// 封面图
func (r *AlibabaalihousenewhomevrsyncAPIRequest) SetCoverImage(_coverImage string) error {
	r._coverImage = _coverImage
	r.Set("cover_image", _coverImage)
	return nil
}

// GetCoverImage CoverImage Getter
func (r AlibabaalihousenewhomevrsyncAPIRequest) GetCoverImage() string {
	return r._coverImage
}

// SetVrUrl is VrUrl Setter
// vr展示链接
func (r *AlibabaalihousenewhomevrsyncAPIRequest) SetVrUrl(_vrUrl string) error {
	r._vrUrl = _vrUrl
	r.Set("vr_url", _vrUrl)
	return nil
}

// GetVrUrl VrUrl Getter
func (r AlibabaalihousenewhomevrsyncAPIRequest) GetVrUrl() string {
	return r._vrUrl
}

// SetMarkingLayoutImg is MarkingLayoutImg Setter
// 营销户型图
func (r *AlibabaalihousenewhomevrsyncAPIRequest) SetMarkingLayoutImg(_markingLayoutImg string) error {
	r._markingLayoutImg = _markingLayoutImg
	r.Set("marking_layout_img", _markingLayoutImg)
	return nil
}

// GetMarkingLayoutImg MarkingLayoutImg Getter
func (r AlibabaalihousenewhomevrsyncAPIRequest) GetMarkingLayoutImg() string {
	return r._markingLayoutImg
}

// SetIsSingleLayout is IsSingleLayout Setter
// 是否是单层 1 是 0 否
func (r *AlibabaalihousenewhomevrsyncAPIRequest) SetIsSingleLayout(_isSingleLayout int64) error {
	r._isSingleLayout = _isSingleLayout
	r.Set("is_single_layout", _isSingleLayout)
	return nil
}

// GetIsSingleLayout IsSingleLayout Getter
func (r AlibabaalihousenewhomevrsyncAPIRequest) GetIsSingleLayout() int64 {
	return r._isSingleLayout
}
