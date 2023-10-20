package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceApplyGetAPIRequest 开票申请数据获取接口 API请求
// alibaba.einvoice.apply.get
//
// ERP获取开票申请数据
type AlibabaEinvoiceApplyGetAPIRequest struct {
	model.Params
	// 平台订单号
	_platformTid string
	// 开票申请ID，跟消息中的apply_id对应，传入applyId后，只会返回一条开票申请消息
	_applyId string
}

// NewAlibabaEinvoiceApplyGetRequest 初始化AlibabaEinvoiceApplyGetAPIRequest对象
func NewAlibabaEinvoiceApplyGetRequest() *AlibabaEinvoiceApplyGetAPIRequest {
	return &AlibabaEinvoiceApplyGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceApplyGetAPIRequest) Reset() {
	r._platformTid = ""
	r._applyId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceApplyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.apply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceApplyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceApplyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlatformTid is PlatformTid Setter
// 平台订单号
func (r *AlibabaEinvoiceApplyGetAPIRequest) SetPlatformTid(_platformTid string) error {
	r._platformTid = _platformTid
	r.Set("platform_tid", _platformTid)
	return nil
}

// GetPlatformTid PlatformTid Getter
func (r AlibabaEinvoiceApplyGetAPIRequest) GetPlatformTid() string {
	return r._platformTid
}

// SetApplyId is ApplyId Setter
// 开票申请ID，跟消息中的apply_id对应，传入applyId后，只会返回一条开票申请消息
func (r *AlibabaEinvoiceApplyGetAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlibabaEinvoiceApplyGetAPIRequest) GetApplyId() string {
	return r._applyId
}

var poolAlibabaEinvoiceApplyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceApplyGetRequest()
	},
}

// GetAlibabaEinvoiceApplyGetRequest 从 sync.Pool 获取 AlibabaEinvoiceApplyGetAPIRequest
func GetAlibabaEinvoiceApplyGetAPIRequest() *AlibabaEinvoiceApplyGetAPIRequest {
	return poolAlibabaEinvoiceApplyGetAPIRequest.Get().(*AlibabaEinvoiceApplyGetAPIRequest)
}

// ReleaseAlibabaEinvoiceApplyGetAPIRequest 将 AlibabaEinvoiceApplyGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceApplyGetAPIRequest(v *AlibabaEinvoiceApplyGetAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceApplyGetAPIRequest.Put(v)
}
