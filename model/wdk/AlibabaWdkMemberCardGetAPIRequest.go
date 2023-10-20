package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmembercardgetAPIRequest 查询会员卡信息 API请求
// alibaba.wdk.member.card.get
//
// 根据会员卡查询会员信息
type AlibabawdkmembercardgetAPIRequest struct {
	model.Params
	// 系统自动生成
	_memberQuery *MemberQueryRequest
}

// NewAlibabawdkmembercardgetRequest 初始化AlibabawdkmembercardgetAPIRequest对象
func NewAlibabawdkmembercardgetRequest() *AlibabawdkmembercardgetAPIRequest {
	return &AlibabawdkmembercardgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmembercardgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.member.card.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmembercardgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmembercardgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemberQuery is MemberQuery Setter
// 系统自动生成
func (r *AlibabawdkmembercardgetAPIRequest) SetMemberQuery(_memberQuery *MemberQueryRequest) error {
	r._memberQuery = _memberQuery
	r.Set("member_query", _memberQuery)
	return nil
}

// GetMemberQuery MemberQuery Getter
func (r AlibabawdkmembercardgetAPIRequest) GetMemberQuery() *MemberQueryRequest {
	return r._memberQuery
}
