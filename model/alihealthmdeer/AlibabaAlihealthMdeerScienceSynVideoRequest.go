package alihealthmdeer

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
视频同步【保存/更新】 API请求
alibaba.alihealth.mdeer.science.synVideo

视频同步【保存/更新】
*/
type AlibabaAlihealthMdeerScienceSynVideoAPIRequest struct {
    model.Params
    // 视频信息实体
    _synVideoInfo   *SynVideoInfo
}

// 初始化AlibabaAlihealthMdeerScienceSynVideoAPIRequest对象
func NewAlibabaAlihealthMdeerScienceSynVideoRequest() *AlibabaAlihealthMdeerScienceSynVideoAPIRequest{
    return &AlibabaAlihealthMdeerScienceSynVideoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMdeerScienceSynVideoAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.mdeer.science.synVideo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMdeerScienceSynVideoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SynVideoInfo Setter
// 视频信息实体
func (r *AlibabaAlihealthMdeerScienceSynVideoAPIRequest) SetSynVideoInfo(_synVideoInfo *SynVideoInfo) error {
    r._synVideoInfo = _synVideoInfo
    r.Set("syn_video_info", _synVideoInfo)
    return nil
}

// SynVideoInfo Getter
func (r AlibabaAlihealthMdeerScienceSynVideoAPIRequest) GetSynVideoInfo() *SynVideoInfo {
    return r._synVideoInfo
}
