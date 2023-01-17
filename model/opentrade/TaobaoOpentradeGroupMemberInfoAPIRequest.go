package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeGroupMemberInfoAPIRequest 组团购获取用户参团信息 API请求
// taobao.opentrade.group.member.info
//
// 组团购场景下，获取用户参团信息
type TaobaoOpentradeGroupMemberInfoAPIRequest struct {
	model.Params
	// 用户openId
	_openUserId string
	// 团id
	_groupId int64
}

// NewTaobaoOpentradeGroupMemberInfoRequest 初始化TaobaoOpentradeGroupMemberInfoAPIRequest对象
func NewTaobaoOpentradeGroupMemberInfoRequest() *TaobaoOpentradeGroupMemberInfoAPIRequest {
	return &TaobaoOpentradeGroupMemberInfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupMemberInfoAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.group.member.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupMemberInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpentradeGroupMemberInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenUserId is OpenUserId Setter
// 用户openId
func (r *TaobaoOpentradeGroupMemberInfoAPIRequest) SetOpenUserId(_openUserId string) error {
	r._openUserId = _openUserId
	r.Set("open_user_id", _openUserId)
	return nil
}

// GetOpenUserId OpenUserId Getter
func (r TaobaoOpentradeGroupMemberInfoAPIRequest) GetOpenUserId() string {
	return r._openUserId
}

// SetGroupId is GroupId Setter
// 团id
func (r *TaobaoOpentradeGroupMemberInfoAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r TaobaoOpentradeGroupMemberInfoAPIRequest) GetGroupId() int64 {
	return r._groupId
}
