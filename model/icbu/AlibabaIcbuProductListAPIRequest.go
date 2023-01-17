package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductListAPIRequest 商品查询 API请求
// alibaba.icbu.product.list
//
// 根据类目ID和商品名称查询商品概要信息。结果以修改时间倒序返回，支持分页，每页最多30个。每次调用都是独立的请求，不记录调用的上下文。
type AlibabaIcbuProductListAPIRequest struct {
	model.Params
	// 商品名称，支持模糊匹配
	_subject string
	// 商品语种，目前只支持ENGLISH
	_language string
	// 最晚修改时间，格式yyyy-MM-dd HH:mm:ss
	_gmtModifiedTo string
	// 最早修改时间，格式yyyy-MM-dd HH:mm:ss
	_gmtModifiedFrom string
	// 类目ID
	_categoryId int64
	// 当前页
	_currentPage int64
	// 每页大小，最大30
	_pageSize int64
	// 商品三级分组id，可选填。若填写-1 则表示取回的商品没有三级分组，不填入代表取回的商品不关心它的三级分组，填写对应的group id将返回这个分组下的商品
	_groupId3 int64
	// 商品二级分组id，可选填。若填写-1 则表示取回的商品没有二级分组，不填入代表取回的商品不关系它的二级分组，填写对应的group id将返回这个分组下的商品
	_groupId2 int64
	// 商品一级分组id，可选填。若填写0 则表示取回的商品没有一级分组，不填入代表取回的商品不关心它的一级分组，填写对应的group id将返回这个分组下的商品
	_groupId1 int64
	// 商品明文id
	_id int64
}

// NewAlibabaIcbuProductListRequest 初始化AlibabaIcbuProductListAPIRequest对象
func NewAlibabaIcbuProductListRequest() *AlibabaIcbuProductListAPIRequest {
	return &AlibabaIcbuProductListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductListAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubject is Subject Setter
// 商品名称，支持模糊匹配
func (r *AlibabaIcbuProductListAPIRequest) SetSubject(_subject string) error {
	r._subject = _subject
	r.Set("subject", _subject)
	return nil
}

// GetSubject Subject Getter
func (r AlibabaIcbuProductListAPIRequest) GetSubject() string {
	return r._subject
}

// SetLanguage is Language Setter
// 商品语种，目前只支持ENGLISH
func (r *AlibabaIcbuProductListAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaIcbuProductListAPIRequest) GetLanguage() string {
	return r._language
}

// SetGmtModifiedTo is GmtModifiedTo Setter
// 最晚修改时间，格式yyyy-MM-dd HH:mm:ss
func (r *AlibabaIcbuProductListAPIRequest) SetGmtModifiedTo(_gmtModifiedTo string) error {
	r._gmtModifiedTo = _gmtModifiedTo
	r.Set("gmt_modified_to", _gmtModifiedTo)
	return nil
}

// GetGmtModifiedTo GmtModifiedTo Getter
func (r AlibabaIcbuProductListAPIRequest) GetGmtModifiedTo() string {
	return r._gmtModifiedTo
}

// SetGmtModifiedFrom is GmtModifiedFrom Setter
// 最早修改时间，格式yyyy-MM-dd HH:mm:ss
func (r *AlibabaIcbuProductListAPIRequest) SetGmtModifiedFrom(_gmtModifiedFrom string) error {
	r._gmtModifiedFrom = _gmtModifiedFrom
	r.Set("gmt_modified_from", _gmtModifiedFrom)
	return nil
}

// GetGmtModifiedFrom GmtModifiedFrom Getter
func (r AlibabaIcbuProductListAPIRequest) GetGmtModifiedFrom() string {
	return r._gmtModifiedFrom
}

// SetCategoryId is CategoryId Setter
// 类目ID
func (r *AlibabaIcbuProductListAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r AlibabaIcbuProductListAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *AlibabaIcbuProductListAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaIcbuProductListAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页大小，最大30
func (r *AlibabaIcbuProductListAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaIcbuProductListAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetGroupId3 is GroupId3 Setter
// 商品三级分组id，可选填。若填写-1 则表示取回的商品没有三级分组，不填入代表取回的商品不关心它的三级分组，填写对应的group id将返回这个分组下的商品
func (r *AlibabaIcbuProductListAPIRequest) SetGroupId3(_groupId3 int64) error {
	r._groupId3 = _groupId3
	r.Set("group_id3", _groupId3)
	return nil
}

// GetGroupId3 GroupId3 Getter
func (r AlibabaIcbuProductListAPIRequest) GetGroupId3() int64 {
	return r._groupId3
}

// SetGroupId2 is GroupId2 Setter
// 商品二级分组id，可选填。若填写-1 则表示取回的商品没有二级分组，不填入代表取回的商品不关系它的二级分组，填写对应的group id将返回这个分组下的商品
func (r *AlibabaIcbuProductListAPIRequest) SetGroupId2(_groupId2 int64) error {
	r._groupId2 = _groupId2
	r.Set("group_id2", _groupId2)
	return nil
}

// GetGroupId2 GroupId2 Getter
func (r AlibabaIcbuProductListAPIRequest) GetGroupId2() int64 {
	return r._groupId2
}

// SetGroupId1 is GroupId1 Setter
// 商品一级分组id，可选填。若填写0 则表示取回的商品没有一级分组，不填入代表取回的商品不关心它的一级分组，填写对应的group id将返回这个分组下的商品
func (r *AlibabaIcbuProductListAPIRequest) SetGroupId1(_groupId1 int64) error {
	r._groupId1 = _groupId1
	r.Set("group_id1", _groupId1)
	return nil
}

// GetGroupId1 GroupId1 Getter
func (r AlibabaIcbuProductListAPIRequest) GetGroupId1() int64 {
	return r._groupId1
}

// SetId is Id Setter
// 商品明文id
func (r *AlibabaIcbuProductListAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaIcbuProductListAPIRequest) GetId() int64 {
	return r._id
}
