package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceapplygetAPIRequest 开票申请数据获取接口 API请求
// alibaba.einvoice.apply.get
//
// ERP获取开票申请数据
type AlibabaeinvoiceapplygetAPIRequest struct {
	model.Params
	// 平台订单号
	_platformTid string
	// 开票申请ID，跟消息中的apply_id对应，传入applyId后，只会返回一条开票申请消息
	_applyId string
}

// NewAlibabaeinvoiceapplygetRequest 初始化AlibabaeinvoiceapplygetAPIRequest对象
func NewAlibabaeinvoiceapplygetRequest() *AlibabaeinvoiceapplygetAPIRequest {
	return &AlibabaeinvoiceapplygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceapplygetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.apply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceapplygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceapplygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlatformTid is PlatformTid Setter
// 平台订单号
func (r *AlibabaeinvoiceapplygetAPIRequest) SetPlatformTid(_platformTid string) error {
	r._platformTid = _platformTid
	r.Set("platform_tid", _platformTid)
	return nil
}

// GetPlatformTid PlatformTid Getter
func (r AlibabaeinvoiceapplygetAPIRequest) GetPlatformTid() string {
	return r._platformTid
}

// SetApplyId is ApplyId Setter
// 开票申请ID，跟消息中的apply_id对应，传入applyId后，只会返回一条开票申请消息
func (r *AlibabaeinvoiceapplygetAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlibabaeinvoiceapplygetAPIRequest) GetApplyId() string {
	return r._applyId
}
