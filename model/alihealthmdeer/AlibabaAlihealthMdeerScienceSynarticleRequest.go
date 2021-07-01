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
type AlibabaAlihealthMdeerScienceSynarticleAPIRequest struct {
    model.Params
    // 同步文章对象
    _synArticleInfo   *SynArticleInfo
}

// 初始化AlibabaAlihealthMdeerScienceSynarticleAPIRequest对象
func NewAlibabaAlihealthMdeerScienceSynarticleRequest() *AlibabaAlihealthMdeerScienceSynarticleAPIRequest{
    return &AlibabaAlihealthMdeerScienceSynarticleAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMdeerScienceSynarticleAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.mdeer.science.synarticle"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMdeerScienceSynarticleAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SynArticleInfo Setter
// 同步文章对象
func (r *AlibabaAlihealthMdeerScienceSynarticleAPIRequest) SetSynArticleInfo(_synArticleInfo *SynArticleInfo) error {
    r._synArticleInfo = _synArticleInfo
    r.Set("syn_article_info", _synArticleInfo)
    return nil
}

// SynArticleInfo Getter
func (r AlibabaAlihealthMdeerScienceSynarticleAPIRequest) GetSynArticleInfo() *SynArticleInfo {
    return r._synArticleInfo
}
