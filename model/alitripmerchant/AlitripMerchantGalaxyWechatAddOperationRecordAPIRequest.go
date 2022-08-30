package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest 用户领取会员卡记录接口 API请求
// alitrip.merchant.galaxy.wechat.add.operation.record
//
// 用户领取会员卡记录接口
type AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户登录凭证
	_token string
	// 添加标识
	_memberCard *MemberCardDto
}

// NewAlitripMerchantGalaxyWechatAddOperationRecordRequest 初始化AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest对象
func NewAlitripMerchantGalaxyWechatAddOperationRecordRequest() *AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest {
	return &AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.add.operation.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录凭证
func (r *AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest) GetToken() string {
	return r._token
}

// SetMemberCard is MemberCard Setter
// 添加标识
func (r *AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest) SetMemberCard(_memberCard *MemberCardDto) error {
	r._memberCard = _memberCard
	r.Set("member_card", _memberCard)
	return nil
}

// GetMemberCard MemberCard Getter
func (r AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest) GetMemberCard() *MemberCardDto {
	return r._memberCard
}
