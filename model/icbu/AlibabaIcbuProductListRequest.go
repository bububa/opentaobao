package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品查询 API请求
alibaba.icbu.product.list

根据类目ID和商品名称查询商品概要信息。结果以修改时间倒序返回，支持分页，每页最多30个。每次调用都是独立的请求，不记录调用的上下文。
*/
type AlibabaIcbuProductListRequest struct {
    model.Params
    // 类目ID
    _categoryId   int64
    // 商品名称，支持模糊匹配
    _subject   string
    // 当前页
    _currentPage   int64
    // 每页大小，最大30
    _pageSize   int64
    // 商品语种，目前只支持ENGLISH
    _language   string
    // 商品三级分组id，可选填。若填写-1 则表示取回的商品没有三级分组，不填入代表取回的商品不关心它的三级分组，填写对应的group id将返回这个分组下的商品
    _groupId3   int64
    // 商品二级分组id，可选填。若填写-1 则表示取回的商品没有二级分组，不填入代表取回的商品不关系它的二级分组，填写对应的group id将返回这个分组下的商品
    _groupId2   int64
    // 商品一级分组id，可选填。若填写0 则表示取回的商品没有一级分组，不填入代表取回的商品不关心它的一级分组，填写对应的group id将返回这个分组下的商品
    _groupId1   int64
    // 商品明文id
    _id   int64
    // 最晚修改时间，格式yyyy-MM-dd HH:mm:ss
    _gmtModifiedTo   string
    // 最早修改时间，格式yyyy-MM-dd HH:mm:ss
    _gmtModifiedFrom   string
}

// 初始化AlibabaIcbuProductListRequest对象
func NewAlibabaIcbuProductListRequest() *AlibabaIcbuProductListRequest{
    return &AlibabaIcbuProductListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductListRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CategoryId Setter
// 类目ID
func (r *AlibabaIcbuProductListRequest) SetCategoryId(_categoryId int64) error {
    r._categoryId = _categoryId
    r.Set("category_id", _categoryId)
    return nil
}

// CategoryId Getter
func (r AlibabaIcbuProductListRequest) GetCategoryId() int64 {
    return r._categoryId
}
// Subject Setter
// 商品名称，支持模糊匹配
func (r *AlibabaIcbuProductListRequest) SetSubject(_subject string) error {
    r._subject = _subject
    r.Set("subject", _subject)
    return nil
}

// Subject Getter
func (r AlibabaIcbuProductListRequest) GetSubject() string {
    return r._subject
}
// CurrentPage Setter
// 当前页
func (r *AlibabaIcbuProductListRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaIcbuProductListRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// PageSize Setter
// 每页大小，最大30
func (r *AlibabaIcbuProductListRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaIcbuProductListRequest) GetPageSize() int64 {
    return r._pageSize
}
// Language Setter
// 商品语种，目前只支持ENGLISH
func (r *AlibabaIcbuProductListRequest) SetLanguage(_language string) error {
    r._language = _language
    r.Set("language", _language)
    return nil
}

// Language Getter
func (r AlibabaIcbuProductListRequest) GetLanguage() string {
    return r._language
}
// GroupId3 Setter
// 商品三级分组id，可选填。若填写-1 则表示取回的商品没有三级分组，不填入代表取回的商品不关心它的三级分组，填写对应的group id将返回这个分组下的商品
func (r *AlibabaIcbuProductListRequest) SetGroupId3(_groupId3 int64) error {
    r._groupId3 = _groupId3
    r.Set("group_id3", _groupId3)
    return nil
}

// GroupId3 Getter
func (r AlibabaIcbuProductListRequest) GetGroupId3() int64 {
    return r._groupId3
}
// GroupId2 Setter
// 商品二级分组id，可选填。若填写-1 则表示取回的商品没有二级分组，不填入代表取回的商品不关系它的二级分组，填写对应的group id将返回这个分组下的商品
func (r *AlibabaIcbuProductListRequest) SetGroupId2(_groupId2 int64) error {
    r._groupId2 = _groupId2
    r.Set("group_id2", _groupId2)
    return nil
}

// GroupId2 Getter
func (r AlibabaIcbuProductListRequest) GetGroupId2() int64 {
    return r._groupId2
}
// GroupId1 Setter
// 商品一级分组id，可选填。若填写0 则表示取回的商品没有一级分组，不填入代表取回的商品不关心它的一级分组，填写对应的group id将返回这个分组下的商品
func (r *AlibabaIcbuProductListRequest) SetGroupId1(_groupId1 int64) error {
    r._groupId1 = _groupId1
    r.Set("group_id1", _groupId1)
    return nil
}

// GroupId1 Getter
func (r AlibabaIcbuProductListRequest) GetGroupId1() int64 {
    return r._groupId1
}
// Id Setter
// 商品明文id
func (r *AlibabaIcbuProductListRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaIcbuProductListRequest) GetId() int64 {
    return r._id
}
// GmtModifiedTo Setter
// 最晚修改时间，格式yyyy-MM-dd HH:mm:ss
func (r *AlibabaIcbuProductListRequest) SetGmtModifiedTo(_gmtModifiedTo string) error {
    r._gmtModifiedTo = _gmtModifiedTo
    r.Set("gmt_modified_to", _gmtModifiedTo)
    return nil
}

// GmtModifiedTo Getter
func (r AlibabaIcbuProductListRequest) GetGmtModifiedTo() string {
    return r._gmtModifiedTo
}
// GmtModifiedFrom Setter
// 最早修改时间，格式yyyy-MM-dd HH:mm:ss
func (r *AlibabaIcbuProductListRequest) SetGmtModifiedFrom(_gmtModifiedFrom string) error {
    r._gmtModifiedFrom = _gmtModifiedFrom
    r.Set("gmt_modified_from", _gmtModifiedFrom)
    return nil
}

// GmtModifiedFrom Getter
func (r AlibabaIcbuProductListRequest) GetGmtModifiedFrom() string {
    return r._gmtModifiedFrom
}
