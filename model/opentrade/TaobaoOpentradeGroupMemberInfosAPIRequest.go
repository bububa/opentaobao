package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeGroupMemberInfosAPIRequest 组团购批量获取用户参团信息 API请求
// taobao.opentrade.group.member.infos
//
// 组团购场景下，获取用户参团信息
type TaobaoOpentradeGroupMemberInfosAPIRequest struct {
	model.Params
	// 用户openId列表
	_openUserIds []string
	// 团id
	_groupId int64
}

// NewTaobaoOpentradeGroupMemberInfosRequest 初始化TaobaoOpentradeGroupMemberInfosAPIRequest对象
func NewTaobaoOpentradeGroupMemberInfosRequest() *TaobaoOpentradeGroupMemberInfosAPIRequest {
	return &TaobaoOpentradeGroupMemberInfosAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupMemberInfosAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.group.member.infos"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupMemberInfosAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOpenUserIds is OpenUserIds Setter
// 用户openId列表
func (r *TaobaoOpentradeGroupMemberInfosAPIRequest) SetOpenUserIds(_openUserIds []string) error {
	r._openUserIds = _openUserIds
	r.Set("open_user_ids", _openUserIds)
	return nil
}

// GetOpenUserIds OpenUserIds Getter
func (r TaobaoOpentradeGroupMemberInfosAPIRequest) GetOpenUserIds() []string {
	return r._openUserIds
}

// SetGroupId is GroupId Setter
// 团id
func (r *TaobaoOpentradeGroupMemberInfosAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r TaobaoOpentradeGroupMemberInfosAPIRequest) GetGroupId() int64 {
	return r._groupId
}
