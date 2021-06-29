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
    title   string
    // 主要内容
    desc   string
    // 图标链接，可以为空
    iconLink   string
    // SHORT/MEDIUM/LONG
    bizType   string
    // 业务类型
    source   string
    // 生效时间，可以为空
    startTime   int64
    // 左按钮文案，可以为空，默认为"取消"
    leftBtnText   string
    // 左按钮链接，可以为空
    leftBtnLink   string
    // 右按钮文案，可以为空，默认为"确定"
    rightBtnText   string
    // 右按钮链接
    rightBtnLink   string
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
func (r *AlibabaTuikeOfferZhitokenRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetTitle() string {
    return r.title
}
// Desc Setter
// 主要内容
func (r *AlibabaTuikeOfferZhitokenRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

// Desc Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetDesc() string {
    return r.desc
}
// IconLink Setter
// 图标链接，可以为空
func (r *AlibabaTuikeOfferZhitokenRequest) SetIconLink(iconLink string) error {
    r.iconLink = iconLink
    r.Set("icon_link", iconLink)
    return nil
}

// IconLink Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetIconLink() string {
    return r.iconLink
}
// BizType Setter
// SHORT/MEDIUM/LONG
func (r *AlibabaTuikeOfferZhitokenRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetBizType() string {
    return r.bizType
}
// Source Setter
// 业务类型
func (r *AlibabaTuikeOfferZhitokenRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetSource() string {
    return r.source
}
// StartTime Setter
// 生效时间，可以为空
func (r *AlibabaTuikeOfferZhitokenRequest) SetStartTime(startTime int64) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetStartTime() int64 {
    return r.startTime
}
// LeftBtnText Setter
// 左按钮文案，可以为空，默认为"取消"
func (r *AlibabaTuikeOfferZhitokenRequest) SetLeftBtnText(leftBtnText string) error {
    r.leftBtnText = leftBtnText
    r.Set("left_btn_text", leftBtnText)
    return nil
}

// LeftBtnText Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetLeftBtnText() string {
    return r.leftBtnText
}
// LeftBtnLink Setter
// 左按钮链接，可以为空
func (r *AlibabaTuikeOfferZhitokenRequest) SetLeftBtnLink(leftBtnLink string) error {
    r.leftBtnLink = leftBtnLink
    r.Set("left_btn_link", leftBtnLink)
    return nil
}

// LeftBtnLink Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetLeftBtnLink() string {
    return r.leftBtnLink
}
// RightBtnText Setter
// 右按钮文案，可以为空，默认为"确定"
func (r *AlibabaTuikeOfferZhitokenRequest) SetRightBtnText(rightBtnText string) error {
    r.rightBtnText = rightBtnText
    r.Set("right_btn_text", rightBtnText)
    return nil
}

// RightBtnText Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetRightBtnText() string {
    return r.rightBtnText
}
// RightBtnLink Setter
// 右按钮链接
func (r *AlibabaTuikeOfferZhitokenRequest) SetRightBtnLink(rightBtnLink string) error {
    r.rightBtnLink = rightBtnLink
    r.Set("right_btn_link", rightBtnLink)
    return nil
}

// RightBtnLink Getter
func (r AlibabaTuikeOfferZhitokenRequest) GetRightBtnLink() string {
    return r.rightBtnLink
}
