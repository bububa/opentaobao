package alihealthmdeer

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
文章删除 API请求
alibaba.alihealth.mdeer.science.deletearticle

三方同步文章删除
*/
type AlibabaAlihealthMdeerScienceDeletearticleRequest struct {
    model.Params
    // 文章ID
    articleId   int64
}

// 初始化AlibabaAlihealthMdeerScienceDeletearticleRequest对象
func NewAlibabaAlihealthMdeerScienceDeletearticleRequest() *AlibabaAlihealthMdeerScienceDeletearticleRequest{
    return &AlibabaAlihealthMdeerScienceDeletearticleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMdeerScienceDeletearticleRequest) GetApiMethodName() string {
    return "alibaba.alihealth.mdeer.science.deletearticle"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMdeerScienceDeletearticleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ArticleId Setter
// 文章ID
func (r *AlibabaAlihealthMdeerScienceDeletearticleRequest) SetArticleId(articleId int64) error {
    r.articleId = articleId
    r.Set("article_id", articleId)
    return nil
}

// ArticleId Getter
func (r AlibabaAlihealthMdeerScienceDeletearticleRequest) GetArticleId() int64 {
    return r.articleId
}
