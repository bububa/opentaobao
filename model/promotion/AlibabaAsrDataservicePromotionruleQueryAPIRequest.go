package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAsrDataservicePromotionruleQueryAPIRequest 星巴克优惠规则查询 API请求
// alibaba.asr.dataservice.promotionrule.query
//
// 查询优惠规则，例如星巴克查询优惠规则
type AlibabaAsrDataservicePromotionruleQueryAPIRequest struct {
	model.Params
	// 当前页
	_pageNo int64
	// 每页数量
	_pageSize int64
}

// NewAlibabaAsrDataservicePromotionruleQueryRequest 初始化AlibabaAsrDataservicePromotionruleQueryAPIRequest对象
func NewAlibabaAsrDataservicePromotionruleQueryRequest() *AlibabaAsrDataservicePromotionruleQueryAPIRequest {
	return &AlibabaAsrDataservicePromotionruleQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAsrDataservicePromotionruleQueryAPIRequest) Reset() {
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAsrDataservicePromotionruleQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.asr.dataservice.promotionrule.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAsrDataservicePromotionruleQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAsrDataservicePromotionruleQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageNo is PageNo Setter
// 当前页
func (r *AlibabaAsrDataservicePromotionruleQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AlibabaAsrDataservicePromotionruleQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页数量
func (r *AlibabaAsrDataservicePromotionruleQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAsrDataservicePromotionruleQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaAsrDataservicePromotionruleQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAsrDataservicePromotionruleQueryRequest()
	},
}

// GetAlibabaAsrDataservicePromotionruleQueryRequest 从 sync.Pool 获取 AlibabaAsrDataservicePromotionruleQueryAPIRequest
func GetAlibabaAsrDataservicePromotionruleQueryAPIRequest() *AlibabaAsrDataservicePromotionruleQueryAPIRequest {
	return poolAlibabaAsrDataservicePromotionruleQueryAPIRequest.Get().(*AlibabaAsrDataservicePromotionruleQueryAPIRequest)
}

// ReleaseAlibabaAsrDataservicePromotionruleQueryAPIRequest 将 AlibabaAsrDataservicePromotionruleQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAsrDataservicePromotionruleQueryAPIRequest(v *AlibabaAsrDataservicePromotionruleQueryAPIRequest) {
	v.Reset()
	poolAlibabaAsrDataservicePromotionruleQueryAPIRequest.Put(v)
}
