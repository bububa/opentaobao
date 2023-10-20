package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripInvoiceGetAPIRequest 获取用户可用发票列表 API请求
// alitrip.btrip.invoice.get
//
// 差旅申请用户获取可用发票列表
type AlitripBtripInvoiceGetAPIRequest struct {
	model.Params
	// 企业id
	_corpId string
	// 用户id
	_userId string
}

// NewAlitripBtripInvoiceGetRequest 初始化AlitripBtripInvoiceGetAPIRequest对象
func NewAlitripBtripInvoiceGetRequest() *AlitripBtripInvoiceGetAPIRequest {
	return &AlitripBtripInvoiceGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripInvoiceGetAPIRequest) Reset() {
	r._corpId = ""
	r._userId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceGetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.invoice.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripInvoiceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCorpId is CorpId Setter
// 企业id
func (r *AlitripBtripInvoiceGetAPIRequest) SetCorpId(_corpId string) error {
	r._corpId = _corpId
	r.Set("corp_id", _corpId)
	return nil
}

// GetCorpId CorpId Getter
func (r AlitripBtripInvoiceGetAPIRequest) GetCorpId() string {
	return r._corpId
}

// SetUserId is UserId Setter
// 用户id
func (r *AlitripBtripInvoiceGetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlitripBtripInvoiceGetAPIRequest) GetUserId() string {
	return r._userId
}

var poolAlitripBtripInvoiceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripInvoiceGetRequest()
	},
}

// GetAlitripBtripInvoiceGetRequest 从 sync.Pool 获取 AlitripBtripInvoiceGetAPIRequest
func GetAlitripBtripInvoiceGetAPIRequest() *AlitripBtripInvoiceGetAPIRequest {
	return poolAlitripBtripInvoiceGetAPIRequest.Get().(*AlitripBtripInvoiceGetAPIRequest)
}

// ReleaseAlitripBtripInvoiceGetAPIRequest 将 AlitripBtripInvoiceGetAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripInvoiceGetAPIRequest(v *AlitripBtripInvoiceGetAPIRequest) {
	v.Reset()
	poolAlitripBtripInvoiceGetAPIRequest.Put(v)
}
