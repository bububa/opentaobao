package icburfq

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuRfqSearchAPIRequest 查询RFQ API请求
// alibaba.icbu.rfq.search
//
// 用于查询RFQ的信息
type AlibabaIcbuRfqSearchAPIRequest struct {
	model.Params
	// 验证
	_md5key string
	// 查询条件
	_cond *RfqRequestSearchCondDto
}

// NewAlibabaIcbuRfqSearchRequest 初始化AlibabaIcbuRfqSearchAPIRequest对象
func NewAlibabaIcbuRfqSearchRequest() *AlibabaIcbuRfqSearchAPIRequest {
	return &AlibabaIcbuRfqSearchAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuRfqSearchAPIRequest) Reset() {
	r._md5key = ""
	r._cond = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.rfq.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuRfqSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMd5key is Md5key Setter
// 验证
func (r *AlibabaIcbuRfqSearchAPIRequest) SetMd5key(_md5key string) error {
	r._md5key = _md5key
	r.Set("md5key", _md5key)
	return nil
}

// GetMd5key Md5key Getter
func (r AlibabaIcbuRfqSearchAPIRequest) GetMd5key() string {
	return r._md5key
}

// SetCond is Cond Setter
// 查询条件
func (r *AlibabaIcbuRfqSearchAPIRequest) SetCond(_cond *RfqRequestSearchCondDto) error {
	r._cond = _cond
	r.Set("cond", _cond)
	return nil
}

// GetCond Cond Getter
func (r AlibabaIcbuRfqSearchAPIRequest) GetCond() *RfqRequestSearchCondDto {
	return r._cond
}

var poolAlibabaIcbuRfqSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuRfqSearchRequest()
	},
}

// GetAlibabaIcbuRfqSearchRequest 从 sync.Pool 获取 AlibabaIcbuRfqSearchAPIRequest
func GetAlibabaIcbuRfqSearchAPIRequest() *AlibabaIcbuRfqSearchAPIRequest {
	return poolAlibabaIcbuRfqSearchAPIRequest.Get().(*AlibabaIcbuRfqSearchAPIRequest)
}

// ReleaseAlibabaIcbuRfqSearchAPIRequest 将 AlibabaIcbuRfqSearchAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuRfqSearchAPIRequest(v *AlibabaIcbuRfqSearchAPIRequest) {
	v.Reset()
	poolAlibabaIcbuRfqSearchAPIRequest.Put(v)
}
