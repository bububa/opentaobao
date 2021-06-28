package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询两个渠道之间的固定映射关系，不通过算法兜底 APIRequest
alibaba.imap.fixedmapping.query

查询两个渠道之间的固定映射关系，不通过算法兜底
*/
type AlibabaImapFixedmappingQueryRequest struct {
    model.Params

    // 密码
    password   string 

    // 账号
    appName   string 

    // 源渠道ID
    srcChannelId   int64 

    // 目标渠道ID列表
    targetChannelIdList   []int64 

    // 目标渠道ID
    targetCategoryId   int64 

    // 源类目ID
    srcCategoryId   int64 

}

func NewAlibabaImapFixedmappingQueryRequest() *AlibabaImapFixedmappingQueryRequest{
    return &AlibabaImapFixedmappingQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaImapFixedmappingQueryRequest) GetApiMethodName() string {
    return "alibaba.imap.fixedmapping.query"
}

func (r AlibabaImapFixedmappingQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaImapFixedmappingQueryRequest) SetPassword(password string) error {
    r.password = password
    r.Set("password", password)
    return nil
}

func (r AlibabaImapFixedmappingQueryRequest) GetPassword() string {
    return r.password
}

func (r *AlibabaImapFixedmappingQueryRequest) SetAppName(appName string) error {
    r.appName = appName
    r.Set("app_name", appName)
    return nil
}

func (r AlibabaImapFixedmappingQueryRequest) GetAppName() string {
    return r.appName
}

func (r *AlibabaImapFixedmappingQueryRequest) SetSrcChannelId(srcChannelId int64) error {
    r.srcChannelId = srcChannelId
    r.Set("src_channel_id", srcChannelId)
    return nil
}

func (r AlibabaImapFixedmappingQueryRequest) GetSrcChannelId() int64 {
    return r.srcChannelId
}

func (r *AlibabaImapFixedmappingQueryRequest) SetTargetChannelIdList(targetChannelIdList []int64) error {
    r.targetChannelIdList = targetChannelIdList
    r.Set("target_channel_id_list", targetChannelIdList)
    return nil
}

func (r AlibabaImapFixedmappingQueryRequest) GetTargetChannelIdList() []int64 {
    return r.targetChannelIdList
}

func (r *AlibabaImapFixedmappingQueryRequest) SetTargetCategoryId(targetCategoryId int64) error {
    r.targetCategoryId = targetCategoryId
    r.Set("target_category_id", targetCategoryId)
    return nil
}

func (r AlibabaImapFixedmappingQueryRequest) GetTargetCategoryId() int64 {
    return r.targetCategoryId
}

func (r *AlibabaImapFixedmappingQueryRequest) SetSrcCategoryId(srcCategoryId int64) error {
    r.srcCategoryId = srcCategoryId
    r.Set("src_category_id", srcCategoryId)
    return nil
}

func (r AlibabaImapFixedmappingQueryRequest) GetSrcCategoryId() int64 {
    return r.srcCategoryId
}

