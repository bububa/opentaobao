package alihealthmdeer

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
文章删除 APIRequest
alibaba.alihealth.mdeer.science.deletearticle

三方同步文章删除
*/
type AlibabaAlihealthMdeerScienceDeletearticleRequest struct {
    model.Params

    // 文章ID
    articleId   int64 

}

func NewAlibabaAlihealthMdeerScienceDeletearticleRequest() *AlibabaAlihealthMdeerScienceDeletearticleRequest{
    return &AlibabaAlihealthMdeerScienceDeletearticleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMdeerScienceDeletearticleRequest) GetApiMethodName() string {
    return "alibaba.alihealth.mdeer.science.deletearticle"
}

func (r AlibabaAlihealthMdeerScienceDeletearticleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMdeerScienceDeletearticleRequest) SetArticleId(articleId int64) error {
    r.articleId = articleId
    r.Set("article_id", articleId)
    return nil
}

func (r AlibabaAlihealthMdeerScienceDeletearticleRequest) GetArticleId() int64 {
    return r.articleId
}

