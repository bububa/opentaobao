package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimrelationsgetAPIRequest 获取openim账号的聊天关系 API请求
// taobao.openim.relations.get
//
// 获取openim账号的聊天关系
type TaobaoopenimrelationsgetAPIRequest struct {
	model.Params
	// 查询起始日期。格式yyyyMMdd。不得早于一个月
	_begDate string
	// 查询结束日期。格式yyyyMMdd。不得早于一个月
	_endDate string
	// 用户信息
	_user *OpenImUser
}

// NewTaobaoopenimrelationsgetRequest 初始化TaobaoopenimrelationsgetAPIRequest对象
func NewTaobaoopenimrelationsgetRequest() *TaobaoopenimrelationsgetAPIRequest {
	return &TaobaoopenimrelationsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimrelationsgetAPIRequest) GetApiMethodName() string {
	return "taobao.openim.relations.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimrelationsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimrelationsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBegDate is BegDate Setter
// 查询起始日期。格式yyyyMMdd。不得早于一个月
func (r *TaobaoopenimrelationsgetAPIRequest) SetBegDate(_begDate string) error {
	r._begDate = _begDate
	r.Set("beg_date", _begDate)
	return nil
}

// GetBegDate BegDate Getter
func (r TaobaoopenimrelationsgetAPIRequest) GetBegDate() string {
	return r._begDate
}

// SetEndDate is EndDate Setter
// 查询结束日期。格式yyyyMMdd。不得早于一个月
func (r *TaobaoopenimrelationsgetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoopenimrelationsgetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoopenimrelationsgetAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoopenimrelationsgetAPIRequest) GetUser() *OpenImUser {
	return r._user
}
