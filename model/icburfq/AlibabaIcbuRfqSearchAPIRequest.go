package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicburfqsearchAPIRequest 查询RFQ API请求
// alibaba.icbu.rfq.search
//
// 用于查询RFQ的信息
type AlibabaicburfqsearchAPIRequest struct {
	model.Params
	// 验证
	_md5key string
	// 查询条件
	_cond *RfqRequestSearchCondDto
}

// NewAlibabaicburfqsearchRequest 初始化AlibabaicburfqsearchAPIRequest对象
func NewAlibabaicburfqsearchRequest() *AlibabaicburfqsearchAPIRequest {
	return &AlibabaicburfqsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicburfqsearchAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.rfq.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicburfqsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicburfqsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMd5key is Md5key Setter
// 验证
func (r *AlibabaicburfqsearchAPIRequest) SetMd5key(_md5key string) error {
	r._md5key = _md5key
	r.Set("md5key", _md5key)
	return nil
}

// GetMd5key Md5key Getter
func (r AlibabaicburfqsearchAPIRequest) GetMd5key() string {
	return r._md5key
}

// SetCond is Cond Setter
// 查询条件
func (r *AlibabaicburfqsearchAPIRequest) SetCond(_cond *RfqRequestSearchCondDto) error {
	r._cond = _cond
	r.Set("cond", _cond)
	return nil
}

// GetCond Cond Getter
func (r AlibabaicburfqsearchAPIRequest) GetCond() *RfqRequestSearchCondDto {
	return r._cond
}
