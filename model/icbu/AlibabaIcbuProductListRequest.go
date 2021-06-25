package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品查询 APIRequest
alibaba.icbu.product.list

根据类目ID和商品名称查询商品概要信息。结果以修改时间倒序返回，支持分页，每页最多30个。每次调用都是独立的请求，不记录调用的上下文。
*/
type AlibabaIcbuProductListRequest struct {
    model.Params

    // 类目ID
    categoryId   int64 

    // 商品名称，支持模糊匹配
    subject   string 

    // 当前页
    currentPage   int64 

    // 每页大小，最大30
    pageSize   int64 

    // 商品语种，目前只支持ENGLISH
    language   string 

    // 商品三级分组id，可选填。若填写-1 则表示取回的商品没有三级分组，不填入代表取回的商品不关心它的三级分组，填写对应的group id将返回这个分组下的商品
    groupId3   int64 

    // 商品二级分组id，可选填。若填写-1 则表示取回的商品没有二级分组，不填入代表取回的商品不关系它的二级分组，填写对应的group id将返回这个分组下的商品
    groupId2   int64 

    // 商品一级分组id，可选填。若填写0 则表示取回的商品没有一级分组，不填入代表取回的商品不关心它的一级分组，填写对应的group id将返回这个分组下的商品
    groupId1   int64 

    // 商品明文id
    id   int64 

    // 最晚修改时间，格式yyyy-MM-dd HH:mm:ss
    gmtModifiedTo   string 

    // 最早修改时间，格式yyyy-MM-dd HH:mm:ss
    gmtModifiedFrom   string 

}

func NewAlibabaIcbuProductListRequest() *AlibabaIcbuProductListRequest{
    return &AlibabaIcbuProductListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductListRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.list"
}

func (r AlibabaIcbuProductListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductListRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r AlibabaIcbuProductListRequest) GetCategoryId() int64 {
    return r.categoryId
}

func (r *AlibabaIcbuProductListRequest) SetSubject(subject string) error {
    r.subject = subject
    r.Set("subject", subject)
    return nil
}

func (r AlibabaIcbuProductListRequest) GetSubject() string {
    return r.subject
}

func (r *AlibabaIcbuProductListRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlibabaIcbuProductListRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *AlibabaIcbuProductListRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaIcbuProductListRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaIcbuProductListRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

func (r AlibabaIcbuProductListRequest) GetLanguage() string {
    return r.language
}

func (r *AlibabaIcbuProductListRequest) SetGroupId3(groupId3 int64) error {
    r.groupId3 = groupId3
    r.Set("group_id3", groupId3)
    return nil
}

func (r AlibabaIcbuProductListRequest) GetGroupId3() int64 {
    return r.groupId3
}

func (r *AlibabaIcbuProductListRequest) SetGroupId2(groupId2 int64) error {
    r.groupId2 = groupId2
    r.Set("group_id2", groupId2)
    return nil
}

func (r AlibabaIcbuProductListRequest) GetGroupId2() int64 {
    return r.groupId2
}

func (r *AlibabaIcbuProductListRequest) SetGroupId1(groupId1 int64) error {
    r.groupId1 = groupId1
    r.Set("group_id1", groupId1)
    return nil
}

func (r AlibabaIcbuProductListRequest) GetGroupId1() int64 {
    return r.groupId1
}

func (r *AlibabaIcbuProductListRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaIcbuProductListRequest) GetId() int64 {
    return r.id
}

func (r *AlibabaIcbuProductListRequest) SetGmtModifiedTo(gmtModifiedTo string) error {
    r.gmtModifiedTo = gmtModifiedTo
    r.Set("gmt_modified_to", gmtModifiedTo)
    return nil
}

func (r AlibabaIcbuProductListRequest) GetGmtModifiedTo() string {
    return r.gmtModifiedTo
}

func (r *AlibabaIcbuProductListRequest) SetGmtModifiedFrom(gmtModifiedFrom string) error {
    r.gmtModifiedFrom = gmtModifiedFrom
    r.Set("gmt_modified_from", gmtModifiedFrom)
    return nil
}

func (r AlibabaIcbuProductListRequest) GetGmtModifiedFrom() string {
    return r.gmtModifiedFrom
}

