package alihealthmdeer

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
视频同步【保存/更新】 APIRequest
alibaba.alihealth.mdeer.science.synVideo

视频同步【保存/更新】
*/
type AlibabaAlihealthMdeerScienceSynVideoRequest struct {
    model.Params

    // 视频信息实体
    synVideoInfo   *SynVideoInfo 

}

func NewAlibabaAlihealthMdeerScienceSynVideoRequest() *AlibabaAlihealthMdeerScienceSynVideoRequest{
    return &AlibabaAlihealthMdeerScienceSynVideoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMdeerScienceSynVideoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.mdeer.science.synVideo"
}

func (r AlibabaAlihealthMdeerScienceSynVideoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMdeerScienceSynVideoRequest) SetSynVideoInfo(synVideoInfo *SynVideoInfo) error {
    r.synVideoInfo = synVideoInfo
    r.Set("syn_video_info", synVideoInfo)
    return nil
}

func (r AlibabaAlihealthMdeerScienceSynVideoRequest) GetSynVideoInfo() *SynVideoInfo {
    return r.synVideoInfo
}

