package alihealthmdeer

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
医知鹿文章同步【保存/更新】 API请求
alibaba.alihealth.mdeer.science.synarticle

文章同步【保存/更新】
*/
type AlibabaAlihealthMdeerScienceSynarticleRequest struct {
    model.Params
    // 同步文章对象
    _synArticleInfo   *SynArticleInfo
}

// 初始化AlibabaAlihealthMdeerScienceSynarticleRequest对象
func NewAlibabaAlihealthMdeerScienceSynarticleRequest() *AlibabaAlihealthMdeerScienceSynarticleRequest{
    return &AlibabaAlihealthMdeerScienceSynarticleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMdeerScienceSynarticleRequest) GetApiMethodName() string {
    return "alibaba.alihealth.mdeer.science.synarticle"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMdeerScienceSynarticleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SynArticleInfo Setter
// 同步文章对象
func (r *AlibabaAlihealthMdeerScienceSynarticleRequest) SetSynArticleInfo(_synArticleInfo *SynArticleInfo) error {
    r._synArticleInfo = _synArticleInfo
    r.Set("syn_article_info", _synArticleInfo)
    return nil
}

// SynArticleInfo Getter
func (r AlibabaAlihealthMdeerScienceSynarticleRequest) GetSynArticleInfo() *SynArticleInfo {
    return r._synArticleInfo
}
