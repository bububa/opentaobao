package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotagTagFindAPIRequest 查询标签接口 API请求
// tmall.promotag.tag.find
//
// 查询用户创建的所有标签
type TmallPromotagTagFindAPIRequest struct {
	model.Params
	// 标签名称，查询时可选项
	_tagName string
	// 每页显示个数
	_pageSize int64
	// 当前页码
	_pageNo int64
	// 标签ID
	_tagId int64
}

// NewTmallPromotagTagFindRequest 初始化TmallPromotagTagFindAPIRequest对象
func NewTmallPromotagTagFindRequest() *TmallPromotagTagFindAPIRequest {
	return &TmallPromotagTagFindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallPromotagTagFindAPIRequest) GetApiMethodName() string {
	return "tmall.promotag.tag.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallPromotagTagFindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallPromotagTagFindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagName is TagName Setter
// 标签名称，查询时可选项
func (r *TmallPromotagTagFindAPIRequest) SetTagName(_tagName string) error {
	r._tagName = _tagName
	r.Set("tag_name", _tagName)
	return nil
}

// GetTagName TagName Getter
func (r TmallPromotagTagFindAPIRequest) GetTagName() string {
	return r._tagName
}

// SetPageSize is PageSize Setter
// 每页显示个数
func (r *TmallPromotagTagFindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallPromotagTagFindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 当前页码
func (r *TmallPromotagTagFindAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TmallPromotagTagFindAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetTagId is TagId Setter
// 标签ID
func (r *TmallPromotagTagFindAPIRequest) SetTagId(_tagId int64) error {
	r._tagId = _tagId
	r.Set("tag_id", _tagId)
	return nil
}

// GetTagId TagId Getter
func (r TmallPromotagTagFindAPIRequest) GetTagId() int64 {
	return r._tagId
}
