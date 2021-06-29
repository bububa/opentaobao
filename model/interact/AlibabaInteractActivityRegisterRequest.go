package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV互动应用活动注册服务 API请求
alibaba.interact.activity.register

为支持卖家由ISV互动应用可以在手淘店铺首页透出，提供ISV互动应用创建的活动注册到手淘的服务
*/
type AlibabaInteractActivityRegisterRequest struct {
    model.Params
    // 页面内容链接，仅允许ascii字符
    entryUrl   string
    // 活动ID
    bizId   string
    // 活动描述文字
    description   string
    // 活动结束时间
    endTime   string
    // 活动名称
    name   string
    // 活动封面图片（非必填）
    picture   string
    // 活动开始时间
    startTime   string
    // 是否有有效时间，若为真开始时间和结束时间必填，默认为真
    hasValidTime   bool
}

// 初始化AlibabaInteractActivityRegisterRequest对象
func NewAlibabaInteractActivityRegisterRequest() *AlibabaInteractActivityRegisterRequest{
    return &AlibabaInteractActivityRegisterRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractActivityRegisterRequest) GetApiMethodName() string {
    return "alibaba.interact.activity.register"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractActivityRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntryUrl Setter
// 页面内容链接，仅允许ascii字符
func (r *AlibabaInteractActivityRegisterRequest) SetEntryUrl(entryUrl string) error {
    r.entryUrl = entryUrl
    r.Set("entry_url", entryUrl)
    return nil
}

// EntryUrl Getter
func (r AlibabaInteractActivityRegisterRequest) GetEntryUrl() string {
    return r.entryUrl
}
// BizId Setter
// 活动ID
func (r *AlibabaInteractActivityRegisterRequest) SetBizId(bizId string) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

// BizId Getter
func (r AlibabaInteractActivityRegisterRequest) GetBizId() string {
    return r.bizId
}
// Description Setter
// 活动描述文字
func (r *AlibabaInteractActivityRegisterRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

// Description Getter
func (r AlibabaInteractActivityRegisterRequest) GetDescription() string {
    return r.description
}
// EndTime Setter
// 活动结束时间
func (r *AlibabaInteractActivityRegisterRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r AlibabaInteractActivityRegisterRequest) GetEndTime() string {
    return r.endTime
}
// Name Setter
// 活动名称
func (r *AlibabaInteractActivityRegisterRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlibabaInteractActivityRegisterRequest) GetName() string {
    return r.name
}
// Picture Setter
// 活动封面图片（非必填）
func (r *AlibabaInteractActivityRegisterRequest) SetPicture(picture string) error {
    r.picture = picture
    r.Set("picture", picture)
    return nil
}

// Picture Getter
func (r AlibabaInteractActivityRegisterRequest) GetPicture() string {
    return r.picture
}
// StartTime Setter
// 活动开始时间
func (r *AlibabaInteractActivityRegisterRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r AlibabaInteractActivityRegisterRequest) GetStartTime() string {
    return r.startTime
}
// HasValidTime Setter
// 是否有有效时间，若为真开始时间和结束时间必填，默认为真
func (r *AlibabaInteractActivityRegisterRequest) SetHasValidTime(hasValidTime bool) error {
    r.hasValidTime = hasValidTime
    r.Set("has_valid_time", hasValidTime)
    return nil
}

// HasValidTime Getter
func (r AlibabaInteractActivityRegisterRequest) GetHasValidTime() bool {
    return r.hasValidTime
}
