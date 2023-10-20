package yunosad

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosadauditcreativeaddAPIRequest 单个创意预审接口 API请求
// yunos.ad.audit.creative.add
//
// YunOS广告业务第三方DSP单个创意预审接口
type YunosadauditcreativeaddAPIRequest struct {
	model.Params
	// 外部dsp的id
	_memberId int64
	// 创意审核入参
	_creative *CreativeParamDto
}

// NewYunosadauditcreativeaddRequest 初始化YunosadauditcreativeaddAPIRequest对象
func NewYunosadauditcreativeaddRequest() *YunosadauditcreativeaddAPIRequest {
	return &YunosadauditcreativeaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosadauditcreativeaddAPIRequest) GetApiMethodName() string {
	return "yunos.ad.audit.creative.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosadauditcreativeaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosadauditcreativeaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemberId is MemberId Setter
// 外部dsp的id
func (r *YunosadauditcreativeaddAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r YunosadauditcreativeaddAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetCreative is Creative Setter
// 创意审核入参
func (r *YunosadauditcreativeaddAPIRequest) SetCreative(_creative *CreativeParamDto) error {
	r._creative = _creative
	r.Set("creative", _creative)
	return nil
}

// GetCreative Creative Getter
func (r YunosadauditcreativeaddAPIRequest) GetCreative() *CreativeParamDto {
	return r._creative
}
