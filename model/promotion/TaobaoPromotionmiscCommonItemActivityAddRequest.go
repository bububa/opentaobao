package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建通用单品优惠活动 API请求
taobao.promotionmisc.common.item.activity.add

创建通用单品优惠活动。
1、该接口只创建活动的基本信息，如需要增加、删除参与该活动的商品及优惠，请调用taobao.promotionmisc.common.item.detail.add和taobao.promotionmisc.common.item.detail.delete接口
2、同一卖家下的活动数量限制为30个，超过限制需先调用taobao.promotionmisc.common.item.activity.delete接口删除无用的活动后才可再创建新的活动
*/
type TaobaoPromotionmiscCommonItemActivityAddRequest struct {
    model.Params
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

// 初始化TaobaoPromotionmiscCommonItemActivityAddRequest对象
func NewTaobaoPromotionmiscCommonItemActivityAddRequest() *TaobaoPromotionmiscCommonItemActivityAddRequest{
    return &TaobaoPromotionmiscCommonItemActivityAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.activity.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 活动名称，不能超过32字符
func (r *TaobaoPromotionmiscCommonItemActivityAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetName() string {
    return r.name
}
// Description Setter
// 活动描述，不能超过100字符
func (r *TaobaoPromotionmiscCommonItemActivityAddRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

// Description Getter
func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetDescription() string {
    return r.description
}
// StartTime Setter
// 活动开始时间
func (r *TaobaoPromotionmiscCommonItemActivityAddRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 活动结束时间
func (r *TaobaoPromotionmiscCommonItemActivityAddRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetEndTime() string {
    return r.endTime
}
// IsUserTag Setter
// 是否指定人群标签
func (r *TaobaoPromotionmiscCommonItemActivityAddRequest) SetIsUserTag(isUserTag bool) error {
    r.isUserTag = isUserTag
    r.Set("is_user_tag", isUserTag)
    return nil
}

// IsUserTag Getter
func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetIsUserTag() bool {
    return r.isUserTag
}
// UserTag Setter
// 用户标签。当is_user_tag为true时，该值才有意义。
func (r *TaobaoPromotionmiscCommonItemActivityAddRequest) SetUserTag(userTag string) error {
    r.userTag = userTag
    r.Set("user_tag", userTag)
    return nil
}

// UserTag Getter
func (r TaobaoPromotionmiscCommonItemActivityAddRequest) GetUserTag() string {
    return r.userTag
}
