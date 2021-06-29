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
    _titleRewriteDetailList   []TitleRewriteDetailDto
    // token来源站点
    _tokenFrom   string
    // 用户token
    _token   string
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
func (r *AlibabaSeakingTitlerewriteSubmitRequest) SetTitleRewriteDetailList(_titleRewriteDetailList []TitleRewriteDetailDto) error {
    r._titleRewriteDetailList = _titleRewriteDetailList
    r.Set("title_rewrite_detail_list", _titleRewriteDetailList)
    return nil
}

// TitleRewriteDetailList Getter
func (r AlibabaSeakingTitlerewriteSubmitRequest) GetTitleRewriteDetailList() []TitleRewriteDetailDto {
    return r._titleRewriteDetailList
}
// TokenFrom Setter
// token来源站点
func (r *AlibabaSeakingTitlerewriteSubmitRequest) SetTokenFrom(_tokenFrom string) error {
    r._tokenFrom = _tokenFrom
    r.Set("token_from", _tokenFrom)
    return nil
}

// TokenFrom Getter
func (r AlibabaSeakingTitlerewriteSubmitRequest) GetTokenFrom() string {
    return r._tokenFrom
}
// Token Setter
// 用户token
func (r *AlibabaSeakingTitlerewriteSubmitRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaSeakingTitlerewriteSubmitRequest) GetToken() string {
    return r._token
}
