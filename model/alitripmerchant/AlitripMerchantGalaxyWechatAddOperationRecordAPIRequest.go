package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxywechataddoperationrecordAPIRequest 用户领取会员卡记录接口 API请求
// alitrip.merchant.galaxy.wechat.add.operation.record
//
// 用户领取会员卡记录接口
type AlitripmerchantgalaxywechataddoperationrecordAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户登录凭证
	_token string
	// 添加标识
	_memberCard *MemberCardDto
}

// NewAlitripmerchantgalaxywechataddoperationrecordRequest 初始化AlitripmerchantgalaxywechataddoperationrecordAPIRequest对象
func NewAlitripmerchantgalaxywechataddoperationrecordRequest() *AlitripmerchantgalaxywechataddoperationrecordAPIRequest {
	return &AlitripmerchantgalaxywechataddoperationrecordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxywechataddoperationrecordAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.add.operation.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxywechataddoperationrecordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxywechataddoperationrecordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripmerchantgalaxywechataddoperationrecordAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxywechataddoperationrecordAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录凭证
func (r *AlitripmerchantgalaxywechataddoperationrecordAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxywechataddoperationrecordAPIRequest) GetToken() string {
	return r._token
}

// SetMemberCard is MemberCard Setter
// 添加标识
func (r *AlitripmerchantgalaxywechataddoperationrecordAPIRequest) SetMemberCard(_memberCard *MemberCardDto) error {
	r._memberCard = _memberCard
	r.Set("member_card", _memberCard)
	return nil
}

// GetMemberCard MemberCard Getter
func (r AlitripmerchantgalaxywechataddoperationrecordAPIRequest) GetMemberCard() *MemberCardDto {
	return r._memberCard
}
