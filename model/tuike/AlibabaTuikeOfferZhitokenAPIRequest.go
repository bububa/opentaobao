package tuike

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatuikeofferzhitokenAPIRequest 生成阿里口令 API请求
// alibaba.tuike.offer.zhitoken
//
// 推荐链接生产吱口令
type AlibabatuikeofferzhitokenAPIRequest struct {
	model.Params
	// 主标题
	_title string
	// 主要内容
	_desc string
	// 图标链接，可以为空
	_iconLink string
	// SHORT/MEDIUM/LONG
	_bizType string
	// 业务类型
	_source string
	// 左按钮文案，可以为空，默认为"取消"
	_leftBtnText string
	// 左按钮链接，可以为空
	_leftBtnLink string
	// 右按钮文案，可以为空，默认为"确定"
	_rightBtnText string
	// 右按钮链接
	_rightBtnLink string
	// 生效时间，可以为空
	_startTime int64
}

// NewAlibabatuikeofferzhitokenRequest 初始化AlibabatuikeofferzhitokenAPIRequest对象
func NewAlibabatuikeofferzhitokenRequest() *AlibabatuikeofferzhitokenAPIRequest {
	return &AlibabatuikeofferzhitokenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatuikeofferzhitokenAPIRequest) GetApiMethodName() string {
	return "alibaba.tuike.offer.zhitoken"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatuikeofferzhitokenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatuikeofferzhitokenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTitle is Title Setter
// 主标题
func (r *AlibabatuikeofferzhitokenAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AlibabatuikeofferzhitokenAPIRequest) GetTitle() string {
	return r._title
}

// SetDesc is Desc Setter
// 主要内容
func (r *AlibabatuikeofferzhitokenAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r AlibabatuikeofferzhitokenAPIRequest) GetDesc() string {
	return r._desc
}

// SetIconLink is IconLink Setter
// 图标链接，可以为空
func (r *AlibabatuikeofferzhitokenAPIRequest) SetIconLink(_iconLink string) error {
	r._iconLink = _iconLink
	r.Set("icon_link", _iconLink)
	return nil
}

// GetIconLink IconLink Getter
func (r AlibabatuikeofferzhitokenAPIRequest) GetIconLink() string {
	return r._iconLink
}

// SetBizType is BizType Setter
// SHORT/MEDIUM/LONG
func (r *AlibabatuikeofferzhitokenAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabatuikeofferzhitokenAPIRequest) GetBizType() string {
	return r._bizType
}

// SetSource is Source Setter
// 业务类型
func (r *AlibabatuikeofferzhitokenAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabatuikeofferzhitokenAPIRequest) GetSource() string {
	return r._source
}

// SetLeftBtnText is LeftBtnText Setter
// 左按钮文案，可以为空，默认为&#34;取消&#34;
func (r *AlibabatuikeofferzhitokenAPIRequest) SetLeftBtnText(_leftBtnText string) error {
	r._leftBtnText = _leftBtnText
	r.Set("left_btn_text", _leftBtnText)
	return nil
}

// GetLeftBtnText LeftBtnText Getter
func (r AlibabatuikeofferzhitokenAPIRequest) GetLeftBtnText() string {
	return r._leftBtnText
}

// SetLeftBtnLink is LeftBtnLink Setter
// 左按钮链接，可以为空
func (r *AlibabatuikeofferzhitokenAPIRequest) SetLeftBtnLink(_leftBtnLink string) error {
	r._leftBtnLink = _leftBtnLink
	r.Set("left_btn_link", _leftBtnLink)
	return nil
}

// GetLeftBtnLink LeftBtnLink Getter
func (r AlibabatuikeofferzhitokenAPIRequest) GetLeftBtnLink() string {
	return r._leftBtnLink
}

// SetRightBtnText is RightBtnText Setter
// 右按钮文案，可以为空，默认为&#34;确定&#34;
func (r *AlibabatuikeofferzhitokenAPIRequest) SetRightBtnText(_rightBtnText string) error {
	r._rightBtnText = _rightBtnText
	r.Set("right_btn_text", _rightBtnText)
	return nil
}

// GetRightBtnText RightBtnText Getter
func (r AlibabatuikeofferzhitokenAPIRequest) GetRightBtnText() string {
	return r._rightBtnText
}

// SetRightBtnLink is RightBtnLink Setter
// 右按钮链接
func (r *AlibabatuikeofferzhitokenAPIRequest) SetRightBtnLink(_rightBtnLink string) error {
	r._rightBtnLink = _rightBtnLink
	r.Set("right_btn_link", _rightBtnLink)
	return nil
}

// GetRightBtnLink RightBtnLink Getter
func (r AlibabatuikeofferzhitokenAPIRequest) GetRightBtnLink() string {
	return r._rightBtnLink
}

// SetStartTime is StartTime Setter
// 生效时间，可以为空
func (r *AlibabatuikeofferzhitokenAPIRequest) SetStartTime(_startTime int64) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AlibabatuikeofferzhitokenAPIRequest) GetStartTime() int64 {
	return r._startTime
}
