package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemClassifyItemQueryAPIRequest
根据分类查商品信息 API请求
taobao.omniitem.classify.item.query

商家根据分类查商品 */
type TaobaoOmniitemClassifyItemQueryAPIRequest struct {
	model.Params
	// 分类ID
	_classifyId int64
	// 页码
	_pageNum int64
	// 每页大小
	_pageSize int64
}

// NewTaobaoOmniitemClassifyItemQueryRequest 初始化TaobaoOmniitemClassifyItemQueryAPIRequest对象
func NewTaobaoOmniitemClassifyItemQueryRequest() *TaobaoOmniitemClassifyItemQueryAPIRequest {
	return &TaobaoOmniitemClassifyItemQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemClassifyItemQueryAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.classify.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemClassifyItemQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ClassifyId Setter
// 分类ID
func (r *TaobaoOmniitemClassifyItemQueryAPIRequest) SetClassifyId(_classifyId int64) error {
	r._classifyId = _classifyId
	r.Set("classify_id", _classifyId)
	return nil
}

// Get ClassifyId Getter
func (r TaobaoOmniitemClassifyItemQueryAPIRequest) GetClassifyId() int64 {
	return r._classifyId
}

// Set is PageNum Setter
// 页码
func (r *TaobaoOmniitemClassifyItemQueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// Get PageNum Getter
func (r TaobaoOmniitemClassifyItemQueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// Set is PageSize Setter
// 每页大小
func (r *TaobaoOmniitemClassifyItemQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoOmniitemClassifyItemQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
