package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeModifytribeinfoAPIRequest
OPENIM群信息修改 API请求
taobao.openim.tribe.modifytribeinfo

OPENIM群信息修改 */
type TaobaoOpenimTribeModifytribeinfoAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群名称
	_tribeName string
	// 群公告
	_notice string
	// 群id
	_tribeId int64
}

// NewTaobaoOpenimTribeModifytribeinfoRequest 初始化TaobaoOpenimTribeModifytribeinfoAPIRequest对象
func NewTaobaoOpenimTribeModifytribeinfoRequest() *TaobaoOpenimTribeModifytribeinfoAPIRequest {
	return &TaobaoOpenimTribeModifytribeinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeModifytribeinfoAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.modifytribeinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeModifytribeinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is User Setter
// 用户信息
func (r *TaobaoOpenimTribeModifytribeinfoAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// Get User Getter
func (r TaobaoOpenimTribeModifytribeinfoAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// Set is TribeName Setter
// 群名称
func (r *TaobaoOpenimTribeModifytribeinfoAPIRequest) SetTribeName(_tribeName string) error {
	r._tribeName = _tribeName
	r.Set("tribe_name", _tribeName)
	return nil
}

// Get TribeName Getter
func (r TaobaoOpenimTribeModifytribeinfoAPIRequest) GetTribeName() string {
	return r._tribeName
}

// Set is Notice Setter
// 群公告
func (r *TaobaoOpenimTribeModifytribeinfoAPIRequest) SetNotice(_notice string) error {
	r._notice = _notice
	r.Set("notice", _notice)
	return nil
}

// Get Notice Getter
func (r TaobaoOpenimTribeModifytribeinfoAPIRequest) GetNotice() string {
	return r._notice
}

// Set is TribeId Setter
// 群id
func (r *TaobaoOpenimTribeModifytribeinfoAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// Get TribeId Getter
func (r TaobaoOpenimTribeModifytribeinfoAPIRequest) GetTribeId() int64 {
	return r._tribeId
}
