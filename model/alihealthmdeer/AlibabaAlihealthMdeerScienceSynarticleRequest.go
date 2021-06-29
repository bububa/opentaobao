package alihealthmdeer

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
医知鹿文章同步【保存/更新】 APIRequest
alibaba.alihealth.mdeer.science.synarticle

文章同步【保存/更新】
*/
type AlibabaAlihealthMdeerScienceSynarticleRequest struct {
    model.Params

    // 同步文章对象
    synArticleInfo   *SynArticleInfo 

}

func NewAlibabaAlihealthMdeerScienceSynarticleRequest() *AlibabaAlihealthMdeerScienceSynarticleRequest{
    return &AlibabaAlihealthMdeerScienceSynarticleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMdeerScienceSynarticleRequest) GetApiMethodName() string {
    return "alibaba.alihealth.mdeer.science.synarticle"
}

func (r AlibabaAlihealthMdeerScienceSynarticleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMdeerScienceSynarticleRequest) SetSynArticleInfo(synArticleInfo *SynArticleInfo) error {
    r.synArticleInfo = synArticleInfo
    r.Set("syn_article_info", synArticleInfo)
    return nil
}

func (r AlibabaAlihealthMdeerScienceSynarticleRequest) GetSynArticleInfo() *SynArticleInfo {
    return r.synArticleInfo
}

