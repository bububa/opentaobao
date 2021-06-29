package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改通用单品优惠活动 API请求
taobao.promotionmisc.common.item.activity.update

修改通用单品优惠活动。
1、该接口只修改活动基本信息，如需要增加、删除参与该活动的商品及优惠，请调用taobao.promotionmisc.common.item.detail.add和taobao.promotionmisc.common.item.detail.delete接口
2、使用该接口时需要把未做修改的字段值也传入
*/
type TaobaoPromotionmiscCommonItemActivityUpdateRequest struct {
    model.Params
    // 优惠活动ID
    activityId   int64
    // 活动名称，不能超过32字符
    name   string
    // 活动描述，不能超过100字符
    description   string
    // 活动开始时间
    startTime   string
    // 活动结束时间
    endTime   string
    // 是否指定人群标签
    isUserTag   bool
    // 用户标签。当is_user_tag为true时，该值才有意义。
    userTag   string
}

// 初始化TaobaoPromotionmiscCommonItemActivityUpdateRequest对象
func NewTaobaoPromotionmiscCommonItemActivityUpdateRequest() *TaobaoPromotionmiscCommonItemActivityUpdateRequest{
    return &TaobaoPromotionmiscCommonItemActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.activity.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 优惠活动ID
func (r *TaobaoPromotionmiscCommonItemActivityUpdateRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetActivityId() int64 {
    return r.activityId
}
// Name Setter
// 活动名称，不能超过32字符
func (r *TaobaoPromotionmiscCommonItemActivityUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetName() string {
    return r.name
}
// Description Setter
// 活动描述，不能超过100字符
func (r *TaobaoPromotionmiscCommonItemActivityUpdateRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

// Description Getter
func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetDescription() string {
    return r.description
}
// StartTime Setter
// 活动开始时间
func (r *TaobaoPromotionmiscCommonItemActivityUpdateRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 活动结束时间
func (r *TaobaoPromotionmiscCommonItemActivityUpdateRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetEndTime() string {
    return r.endTime
}
// IsUserTag Setter
// 是否指定人群标签
func (r *TaobaoPromotionmiscCommonItemActivityUpdateRequest) SetIsUserTag(isUserTag bool) error {
    r.isUserTag = isUserTag
    r.Set("is_user_tag", isUserTag)
    return nil
}

// IsUserTag Getter
func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetIsUserTag() bool {
    return r.isUserTag
}
// UserTag Setter
// 用户标签。当is_user_tag为true时，该值才有意义。
func (r *TaobaoPromotionmiscCommonItemActivityUpdateRequest) SetUserTag(userTag string) error {
    r.userTag = userTag
    r.Set("user_tag", userTag)
    return nil
}

// UserTag Getter
func (r TaobaoPromotionmiscCommonItemActivityUpdateRequest) GetUserTag() string {
    return r.userTag
}
