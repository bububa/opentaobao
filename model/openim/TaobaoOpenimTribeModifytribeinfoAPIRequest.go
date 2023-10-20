package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribemodifytribeinfoAPIRequest OPENIM群信息修改 API请求
// taobao.openim.tribe.modifytribeinfo
//
// OPENIM群信息修改
type TaobaoopenimtribemodifytribeinfoAPIRequest struct {
	model.Params
	// 群名称
	_tribeName string
	// 群公告
	_notice string
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
}

// NewTaobaoopenimtribemodifytribeinfoRequest 初始化TaobaoopenimtribemodifytribeinfoAPIRequest对象
func NewTaobaoopenimtribemodifytribeinfoRequest() *TaobaoopenimtribemodifytribeinfoAPIRequest {
	return &TaobaoopenimtribemodifytribeinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimtribemodifytribeinfoAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.modifytribeinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimtribemodifytribeinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimtribemodifytribeinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTribeName is TribeName Setter
// 群名称
func (r *TaobaoopenimtribemodifytribeinfoAPIRequest) SetTribeName(_tribeName string) error {
	r._tribeName = _tribeName
	r.Set("tribe_name", _tribeName)
	return nil
}

// GetTribeName TribeName Getter
func (r TaobaoopenimtribemodifytribeinfoAPIRequest) GetTribeName() string {
	return r._tribeName
}

// SetNotice is Notice Setter
// 群公告
func (r *TaobaoopenimtribemodifytribeinfoAPIRequest) SetNotice(_notice string) error {
	r._notice = _notice
	r.Set("notice", _notice)
	return nil
}

// GetNotice Notice Getter
func (r TaobaoopenimtribemodifytribeinfoAPIRequest) GetNotice() string {
	return r._notice
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoopenimtribemodifytribeinfoAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoopenimtribemodifytribeinfoAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoopenimtribemodifytribeinfoAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoopenimtribemodifytribeinfoAPIRequest) GetTribeId() int64 {
	return r._tribeId
}
