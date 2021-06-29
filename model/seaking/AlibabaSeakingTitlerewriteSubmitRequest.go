package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交标题改写任务 API请求
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

// 初始化AlibabaSeakingTitlerewriteSubmitRequest对象
func NewAlibabaSeakingTitlerewriteSubmitRequest() *AlibabaSeakingTitlerewriteSubmitRequest{
    return &AlibabaSeakingTitlerewriteSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingTitlerewriteSubmitRequest) GetApiMethodName() string {
    return "alibaba.seaking.titlerewrite.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingTitlerewriteSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TitleRewriteDetailList Setter
// 任务详情列表
func (r *AlibabaSeakingTitlerewriteSubmitRequest) SetTitleRewriteDetailList(titleRewriteDetailList []TitleRewriteDetailDto) error {
    r.titleRewriteDetailList = titleRewriteDetailList
    r.Set("title_rewrite_detail_list", titleRewriteDetailList)
    return nil
}

// TitleRewriteDetailList Getter
func (r AlibabaSeakingTitlerewriteSubmitRequest) GetTitleRewriteDetailList() []TitleRewriteDetailDto {
    return r.titleRewriteDetailList
}
// TokenFrom Setter
// token来源站点
func (r *AlibabaSeakingTitlerewriteSubmitRequest) SetTokenFrom(tokenFrom string) error {
    r.tokenFrom = tokenFrom
    r.Set("token_from", tokenFrom)
    return nil
}

// TokenFrom Getter
func (r AlibabaSeakingTitlerewriteSubmitRequest) GetTokenFrom() string {
    return r.tokenFrom
}
// Token Setter
// 用户token
func (r *AlibabaSeakingTitlerewriteSubmitRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlibabaSeakingTitlerewriteSubmitRequest) GetToken() string {
    return r.token
}
