package icburfq

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.rfq.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
