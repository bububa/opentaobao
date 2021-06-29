package tuike

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
生成阿里口令 API请求
alibaba.tuike.offer.zhitoken

推荐链接生产吱口令
*/
type AlibabaTuikeOfferZhitokenRequest struct {
    model.Params
    // 主标题
    _title   string
    // 主要内容
    _desc   string
    // 图标链接，可以为空
    _iconLink   string
    // SHORT/MEDIUM/LONG
    _bizType   string
    // 业务类型
    _source   string
    // 生效时间，可以为空
    _startTime   int64
    // 左按钮文案，可以为空，默认为"取消"
    _leftBtnText   string
    // 左按钮链接，可以为空
    _leftBtnLink   string
    // 右按钮文案，可以为空，默认为"确定"
    _rightBtnText   string
    // 右按钮链接
    _rightBtnLink   string
}

// 初始化AlibabaTuikeOfferZhitokenRequest对象
func NewAlibabaTuikeOfferZhitokenRequest() *AlibabaTuikeOfferZhitokenRequest{
    return &AlibabaTuikeOfferZhitokenRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTuikeOfferZhitokenRequest) GetApiMethodName() string {
    return "alibaba.tuike.offer.zhitoken"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTuikeOfferZhitokenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Title Setter
// 主标题
func (r *AlibabaTuikeOfferZhitokenRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetTitle() string {
    return r._title
}
// Desc Setter
// 主要内容
func (r *AlibabaTuikeOfferZhitokenRequest) SetDesc(_desc string) error {
    r._desc = _desc
    r.Set("desc", _desc)
    return nil
}

// Desc Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetDesc() string {
    return r._desc
}
// IconLink Setter
// 图标链接，可以为空
func (r *AlibabaTuikeOfferZhitokenRequest) SetIconLink(_iconLink string) error {
    r._iconLink = _iconLink
    r.Set("icon_link", _iconLink)
    return nil
}

// IconLink Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetIconLink() string {
    return r._iconLink
}
// BizType Setter
// SHORT/MEDIUM/LONG
func (r *AlibabaTuikeOfferZhitokenRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetBizType() string {
    return r._bizType
}
// Source Setter
// 业务类型
func (r *AlibabaTuikeOfferZhitokenRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetSource() string {
    return r._source
}
// StartTime Setter
// 生效时间，可以为空
func (r *AlibabaTuikeOfferZhitokenRequest) SetStartTime(_startTime int64) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetStartTime() int64 {
    return r._startTime
}
// LeftBtnText Setter
// 左按钮文案，可以为空，默认为"取消"
func (r *AlibabaTuikeOfferZhitokenRequest) SetLeftBtnText(_leftBtnText string) error {
    r._leftBtnText = _leftBtnText
    r.Set("left_btn_text", _leftBtnText)
    return nil
}

// LeftBtnText Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetLeftBtnText() string {
    return r._leftBtnText
}
// LeftBtnLink Setter
// 左按钮链接，可以为空
func (r *AlibabaTuikeOfferZhitokenRequest) SetLeftBtnLink(_leftBtnLink string) error {
    r._leftBtnLink = _leftBtnLink
    r.Set("left_btn_link", _leftBtnLink)
    return nil
}

// LeftBtnLink Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetLeftBtnLink() string {
    return r._leftBtnLink
}
// RightBtnText Setter
// 右按钮文案，可以为空，默认为"确定"
func (r *AlibabaTuikeOfferZhitokenRequest) SetRightBtnText(_rightBtnText string) error {
    r._rightBtnText = _rightBtnText
    r.Set("right_btn_text", _rightBtnText)
    return nil
}

// RightBtnText Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetRightBtnText() string {
    return r._rightBtnText
}
// RightBtnLink Setter
// 右按钮链接
func (r *AlibabaTuikeOfferZhitokenRequest) SetRightBtnLink(_rightBtnLink string) error {
    r._rightBtnLink = _rightBtnLink
    r.Set("right_btn_link", _rightBtnLink)
    return nil
}

// RightBtnLink Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetRightBtnLink() string {
    return r._rightBtnLink
}
