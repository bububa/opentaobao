package ioti

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitesleslimageshowimagecommonAPIRequest 对混合云提供的刷图接口 API请求
// alibaba.it.esl.eslimage.showimagecommon
//
// 混合云使用，提供给isv和我们混合云环境部署的应用刷图
type AlibabaitesleslimageshowimagecommonAPIRequest struct {
	model.Params
	// ma地址
	_mac string
	// 图片的base64编码,图片要和价签大小一致
	_content2 string
	// 图片2的base64编码,图片要和价签大小一致
	_content string
	// 是否压缩，默认不传，字符串：yes，no
	_isCompress string
	// 是否手动刷图，默认不传，字符串：true，false
	_isManual string
}

// NewAlibabaitesleslimageshowimagecommonRequest 初始化AlibabaitesleslimageshowimagecommonAPIRequest对象
func NewAlibabaitesleslimageshowimagecommonRequest() *AlibabaitesleslimageshowimagecommonAPIRequest {
	return &AlibabaitesleslimageshowimagecommonAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaitesleslimageshowimagecommonAPIRequest) GetApiMethodName() string {
	return "alibaba.it.esl.eslimage.showimagecommon"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaitesleslimageshowimagecommonAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaitesleslimageshowimagecommonAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMac is Mac Setter
// ma地址
func (r *AlibabaitesleslimageshowimagecommonAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// GetMac Mac Getter
func (r AlibabaitesleslimageshowimagecommonAPIRequest) GetMac() string {
	return r._mac
}

// SetContent2 is Content2 Setter
// 图片的base64编码,图片要和价签大小一致
func (r *AlibabaitesleslimageshowimagecommonAPIRequest) SetContent2(_content2 string) error {
	r._content2 = _content2
	r.Set("content2", _content2)
	return nil
}

// GetContent2 Content2 Getter
func (r AlibabaitesleslimageshowimagecommonAPIRequest) GetContent2() string {
	return r._content2
}

// SetContent is Content Setter
// 图片2的base64编码,图片要和价签大小一致
func (r *AlibabaitesleslimageshowimagecommonAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaitesleslimageshowimagecommonAPIRequest) GetContent() string {
	return r._content
}

// SetIsCompress is IsCompress Setter
// 是否压缩，默认不传，字符串：yes，no
func (r *AlibabaitesleslimageshowimagecommonAPIRequest) SetIsCompress(_isCompress string) error {
	r._isCompress = _isCompress
	r.Set("is_compress", _isCompress)
	return nil
}

// GetIsCompress IsCompress Getter
func (r AlibabaitesleslimageshowimagecommonAPIRequest) GetIsCompress() string {
	return r._isCompress
}

// SetIsManual is IsManual Setter
// 是否手动刷图，默认不传，字符串：true，false
func (r *AlibabaitesleslimageshowimagecommonAPIRequest) SetIsManual(_isManual string) error {
	r._isManual = _isManual
	r.Set("is_manual", _isManual)
	return nil
}

// GetIsManual IsManual Getter
func (r AlibabaitesleslimageshowimagecommonAPIRequest) GetIsManual() string {
	return r._isManual
}
