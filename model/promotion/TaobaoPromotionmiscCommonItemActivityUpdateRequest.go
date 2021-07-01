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
type TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest struct {
    model.Params
    // 优惠活动ID
    _activityId   int64
    // 活动名称，不能超过32字符
    _name   string
    // 活动描述，不能超过100字符
    _description   string
    // 活动开始时间
    _startTime   string
    // 活动结束时间
    _endTime   string
    // 是否指定人群标签
    _isUserTag   bool
    // 用户标签。当is_user_tag为true时，该值才有意义。
    _userTag   string
}

// 初始化TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest对象
func NewTaobaoPromotionmiscCommonItemActivityUpdateRequest() *TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest{
    return &TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.activity.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 优惠活动ID
func (r *TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) SetActivityId(_activityId int64) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) GetActivityId() int64 {
    return r._activityId
}
// Name Setter
// 活动名称，不能超过32字符
func (r *TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) GetName() string {
    return r._name
}
// Description Setter
// 活动描述，不能超过100字符
func (r *TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) SetDescription(_description string) error {
    r._description = _description
    r.Set("description", _description)
    return nil
}

// Description Getter
func (r TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) GetDescription() string {
    return r._description
}
// StartTime Setter
// 活动开始时间
func (r *TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 活动结束时间
func (r *TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) GetEndTime() string {
    return r._endTime
}
// IsUserTag Setter
// 是否指定人群标签
func (r *TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) SetIsUserTag(_isUserTag bool) error {
    r._isUserTag = _isUserTag
    r.Set("is_user_tag", _isUserTag)
    return nil
}

// IsUserTag Getter
func (r TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) GetIsUserTag() bool {
    return r._isUserTag
}
// UserTag Setter
// 用户标签。当is_user_tag为true时，该值才有意义。
func (r *TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) SetUserTag(_userTag string) error {
    r._userTag = _userTag
    r.Set("user_tag", _userTag)
    return nil
}

// UserTag Getter
func (r TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest) GetUserTag() string {
    return r._userTag
}
