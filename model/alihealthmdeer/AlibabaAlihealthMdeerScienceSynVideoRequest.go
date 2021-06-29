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
type AlibabaAlihealthMdeerScienceSynVideoRequest struct {
    model.Params
    // 视频信息实体
    _synVideoInfo   *SynVideoInfo
}

// 初始化AlibabaAlihealthMdeerScienceSynVideoRequest对象
func NewAlibabaAlihealthMdeerScienceSynVideoRequest() *AlibabaAlihealthMdeerScienceSynVideoRequest{
    return &AlibabaAlihealthMdeerScienceSynVideoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMdeerScienceSynVideoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.mdeer.science.synVideo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMdeerScienceSynVideoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SynVideoInfo Setter
// 视频信息实体
func (r *AlibabaAlihealthMdeerScienceSynVideoRequest) SetSynVideoInfo(_synVideoInfo *SynVideoInfo) error {
    r._synVideoInfo = _synVideoInfo
    r.Set("syn_video_info", _synVideoInfo)
    return nil
}

// SynVideoInfo Getter
func (r AlibabaAlihealthMdeerScienceSynVideoRequest) GetSynVideoInfo() *SynVideoInfo {
    return r._synVideoInfo
}
