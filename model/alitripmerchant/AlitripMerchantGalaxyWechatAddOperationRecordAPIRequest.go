package alitripmerchant

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._memberCard = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.add.operation.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripMerchantGalaxyWechatAddOperationRecordAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyWechatAddOperationRecordRequest()
	},
}

// GetAlitripMerchantGalaxyWechatAddOperationRecordRequest 从 sync.Pool 获取 AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest
func GetAlitripMerchantGalaxyWechatAddOperationRecordAPIRequest() *AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest {
	return poolAlitripMerchantGalaxyWechatAddOperationRecordAPIRequest.Get().(*AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest)
}

// ReleaseAlitripMerchantGalaxyWechatAddOperationRecordAPIRequest 将 AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatAddOperationRecordAPIRequest(v *AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatAddOperationRecordAPIRequest.Put(v)
}
