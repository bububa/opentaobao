package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAsrDataservicePromotionruleQueryAPIRequest
星巴克优惠规则查询 API请求
alibaba.asr.dataservice.promotionrule.query

查询优惠规则，例如星巴克查询优惠规则 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAsrDataservicePromotionruleQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.asr.dataservice.promotionrule.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAsrDataservicePromotionruleQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PageNo Setter
// 当前页
func (r *AlibabaAsrDataservicePromotionruleQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r AlibabaAsrDataservicePromotionruleQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 每页数量
func (r *AlibabaAsrDataservicePromotionruleQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaAsrDataservicePromotionruleQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
