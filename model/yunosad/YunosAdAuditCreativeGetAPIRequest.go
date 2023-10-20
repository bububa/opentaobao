package yunosad

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosadauditcreativegetAPIRequest 获取单个创意审核状态 API请求
// yunos.ad.audit.creative.get
//
// 获取单个创意审核状态
type YunosadauditcreativegetAPIRequest struct {
	model.Params
	// 第三方广告创意id
	_creativeId string
	// 第三方的dspId
	_memberId int64
}

// NewYunosadauditcreativegetRequest 初始化YunosadauditcreativegetAPIRequest对象
func NewYunosadauditcreativegetRequest() *YunosadauditcreativegetAPIRequest {
	return &YunosadauditcreativegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosadauditcreativegetAPIRequest) GetApiMethodName() string {
	return "yunos.ad.audit.creative.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosadauditcreativegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosadauditcreativegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreativeId is CreativeId Setter
// 第三方广告创意id
func (r *YunosadauditcreativegetAPIRequest) SetCreativeId(_creativeId string) error {
	r._creativeId = _creativeId
	r.Set("creative_id", _creativeId)
	return nil
}

// GetCreativeId CreativeId Getter
func (r YunosadauditcreativegetAPIRequest) GetCreativeId() string {
	return r._creativeId
}

// SetMemberId is MemberId Setter
// 第三方的dspId
func (r *YunosadauditcreativegetAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r YunosadauditcreativegetAPIRequest) GetMemberId() int64 {
	return r._memberId
}
