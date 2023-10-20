package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWholesaleShippinglineTemplateListAPIRequest 获取运费模板 API请求
// alibaba.wholesale.shippingline.template.list
//
// 查询运费模板信息
type AlibabaWholesaleShippinglineTemplateListAPIRequest struct {
	model.Params
	// 第几页从1开始
	_pageNum int64
	// 每页返回的数据个数
	_count int64
}

// NewAlibabaWholesaleShippinglineTemplateListRequest 初始化AlibabaWholesaleShippinglineTemplateListAPIRequest对象
func NewAlibabaWholesaleShippinglineTemplateListRequest() *AlibabaWholesaleShippinglineTemplateListAPIRequest {
	return &AlibabaWholesaleShippinglineTemplateListAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWholesaleShippinglineTemplateListAPIRequest) Reset() {
	r._pageNum = 0
	r._count = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWholesaleShippinglineTemplateListAPIRequest) GetApiMethodName() string {
	return "alibaba.wholesale.shippingline.template.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWholesaleShippinglineTemplateListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWholesaleShippinglineTemplateListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageNum is PageNum Setter
// 第几页从1开始
func (r *AlibabaWholesaleShippinglineTemplateListAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabaWholesaleShippinglineTemplateListAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetCount is Count Setter
// 每页返回的数据个数
func (r *AlibabaWholesaleShippinglineTemplateListAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// GetCount Count Getter
func (r AlibabaWholesaleShippinglineTemplateListAPIRequest) GetCount() int64 {
	return r._count
}

var poolAlibabaWholesaleShippinglineTemplateListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWholesaleShippinglineTemplateListRequest()
	},
}

// GetAlibabaWholesaleShippinglineTemplateListRequest 从 sync.Pool 获取 AlibabaWholesaleShippinglineTemplateListAPIRequest
func GetAlibabaWholesaleShippinglineTemplateListAPIRequest() *AlibabaWholesaleShippinglineTemplateListAPIRequest {
	return poolAlibabaWholesaleShippinglineTemplateListAPIRequest.Get().(*AlibabaWholesaleShippinglineTemplateListAPIRequest)
}

// ReleaseAlibabaWholesaleShippinglineTemplateListAPIRequest 将 AlibabaWholesaleShippinglineTemplateListAPIRequest 放入 sync.Pool
func ReleaseAlibabaWholesaleShippinglineTemplateListAPIRequest(v *AlibabaWholesaleShippinglineTemplateListAPIRequest) {
	v.Reset()
	poolAlibabaWholesaleShippinglineTemplateListAPIRequest.Put(v)
}
