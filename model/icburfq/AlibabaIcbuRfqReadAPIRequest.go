package icburfq

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuRfqReadAPIRequest 是否已读RFQ API请求
// alibaba.icbu.rfq.read
//
// 是否已读RFQ
type AlibabaIcbuRfqReadAPIRequest struct {
	model.Params
	// 查询RFQID列表
	_rfqIdList []string
}

// NewAlibabaIcbuRfqReadRequest 初始化AlibabaIcbuRfqReadAPIRequest对象
func NewAlibabaIcbuRfqReadRequest() *AlibabaIcbuRfqReadAPIRequest {
	return &AlibabaIcbuRfqReadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuRfqReadAPIRequest) Reset() {
	r._rfqIdList = r._rfqIdList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqReadAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.rfq.read"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqReadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuRfqReadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRfqIdList is RfqIdList Setter
// 查询RFQID列表
func (r *AlibabaIcbuRfqReadAPIRequest) SetRfqIdList(_rfqIdList []string) error {
	r._rfqIdList = _rfqIdList
	r.Set("rfq_id_list", _rfqIdList)
	return nil
}

// GetRfqIdList RfqIdList Getter
func (r AlibabaIcbuRfqReadAPIRequest) GetRfqIdList() []string {
	return r._rfqIdList
}

var poolAlibabaIcbuRfqReadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuRfqReadRequest()
	},
}

// GetAlibabaIcbuRfqReadRequest 从 sync.Pool 获取 AlibabaIcbuRfqReadAPIRequest
func GetAlibabaIcbuRfqReadAPIRequest() *AlibabaIcbuRfqReadAPIRequest {
	return poolAlibabaIcbuRfqReadAPIRequest.Get().(*AlibabaIcbuRfqReadAPIRequest)
}

// ReleaseAlibabaIcbuRfqReadAPIRequest 将 AlibabaIcbuRfqReadAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuRfqReadAPIRequest(v *AlibabaIcbuRfqReadAPIRequest) {
	v.Reset()
	poolAlibabaIcbuRfqReadAPIRequest.Put(v)
}
