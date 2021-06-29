package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交标题改写任务 APIRequest
alibaba.seaking.titlerewrite.submit

提交标题改写任务
*/
type AlibabaSeakingTitlerewriteSubmitRequest struct {
    model.Params

    // 任务详情列表
    titleRewriteDetailList   []TitleRewriteDetailDto 

    // token来源站点
    tokenFrom   string 

    // 用户token
    token   string 

}

func NewAlibabaSeakingTitlerewriteSubmitRequest() *AlibabaSeakingTitlerewriteSubmitRequest{
    return &AlibabaSeakingTitlerewriteSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSeakingTitlerewriteSubmitRequest) GetApiMethodName() string {
    return "alibaba.seaking.titlerewrite.submit"
}

func (r AlibabaSeakingTitlerewriteSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSeakingTitlerewriteSubmitRequest) SetTitleRewriteDetailList(titleRewriteDetailList []TitleRewriteDetailDto) error {
    r.titleRewriteDetailList = titleRewriteDetailList
    r.Set("title_rewrite_detail_list", titleRewriteDetailList)
    return nil
}

func (r AlibabaSeakingTitlerewriteSubmitRequest) GetTitleRewriteDetailList() []TitleRewriteDetailDto {
    return r.titleRewriteDetailList
}

func (r *AlibabaSeakingTitlerewriteSubmitRequest) SetTokenFrom(tokenFrom string) error {
    r.tokenFrom = tokenFrom
    r.Set("token_from", tokenFrom)
    return nil
}

func (r AlibabaSeakingTitlerewriteSubmitRequest) GetTokenFrom() string {
    return r.tokenFrom
}

func (r *AlibabaSeakingTitlerewriteSubmitRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaSeakingTitlerewriteSubmitRequest) GetToken() string {
    return r.token
}

