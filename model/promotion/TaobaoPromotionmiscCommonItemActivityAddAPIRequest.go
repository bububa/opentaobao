package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemActivityAddAPIRequest 创建通用单品优惠活动 API请求
// taobao.promotionmisc.common.item.activity.add
//
// 创建通用单品优惠活动。
// 1、该接口只创建活动的基本信息，如需要增加、删除参与该活动的商品及优惠，请调用taobao.promotionmisc.common.item.detail.add和taobao.promotionmisc.common.item.detail.delete接口
// 2、同一卖家下的活动数量限制为30个，超过限制需先调用taobao.promotionmisc.common.item.activity.delete接口删除无用的活动后才可再创建新的活动
type TaobaoPromotionmiscCommonItemActivityAddAPIRequest struct {
	model.Params
	// 活动描述，不能超过100字符
	_description string
	// 活动结束时间
	_endTime string
	// 活动名称，不能超过32字符
	_name string
	// 活动开始时间
	_startTime string
	// 用户标签。当is_user_tag为true时，该值才有意义。
	_userTag string
	// 是否指定人群标签
	_isUserTag bool
}

// NewTaobaoPromotionmiscCommonItemActivityAddRequest 初始化TaobaoPromotionmiscCommonItemActivityAddAPIRequest对象
func NewTaobaoPromotionmiscCommonItemActivityAddRequest() *TaobaoPromotionmiscCommonItemActivityAddAPIRequest {
	return &TaobaoPromotionmiscCommonItemActivityAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemActivityAddAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.activity.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemActivityAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDescription is Description Setter
// 活动描述，不能超过100字符
func (r *TaobaoPromotionmiscCommonItemActivityAddAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r TaobaoPromotionmiscCommonItemActivityAddAPIRequest) GetDescription() string {
	return r._description
}

// SetEndTime is EndTime Setter
// 活动结束时间
func (r *TaobaoPromotionmiscCommonItemActivityAddAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoPromotionmiscCommonItemActivityAddAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetName is Name Setter
// 活动名称，不能超过32字符
func (r *TaobaoPromotionmiscCommonItemActivityAddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoPromotionmiscCommonItemActivityAddAPIRequest) GetName() string {
	return r._name
}

// SetStartTime is StartTime Setter
// 活动开始时间
func (r *TaobaoPromotionmiscCommonItemActivityAddAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoPromotionmiscCommonItemActivityAddAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetUserTag is UserTag Setter
// 用户标签。当is_user_tag为true时，该值才有意义。
func (r *TaobaoPromotionmiscCommonItemActivityAddAPIRequest) SetUserTag(_userTag string) error {
	r._userTag = _userTag
	r.Set("user_tag", _userTag)
	return nil
}

// GetUserTag UserTag Getter
func (r TaobaoPromotionmiscCommonItemActivityAddAPIRequest) GetUserTag() string {
	return r._userTag
}

// SetIsUserTag is IsUserTag Setter
// 是否指定人群标签
func (r *TaobaoPromotionmiscCommonItemActivityAddAPIRequest) SetIsUserTag(_isUserTag bool) error {
	r._isUserTag = _isUserTag
	r.Set("is_user_tag", _isUserTag)
	return nil
}

// GetIsUserTag IsUserTag Getter
func (r TaobaoPromotionmiscCommonItemActivityAddAPIRequest) GetIsUserTag() bool {
	return r._isUserTag
}
