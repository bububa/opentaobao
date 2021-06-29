package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里房产图文草稿信息同步 APIRequest
alibaba.alihouse.newhome.rc.sync

接收图文草稿信息
*/
type AlibabaAlihouseNewhomeRcSyncRequest struct {
    model.Params

    // 图文内容
    rc   *RichContentDraftDto 

}

func NewAlibabaAlihouseNewhomeRcSyncRequest() *AlibabaAlihouseNewhomeRcSyncRequest{
    return &AlibabaAlihouseNewhomeRcSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeRcSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.rc.sync"
}

func (r AlibabaAlihouseNewhomeRcSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeRcSyncRequest) SetRc(rc *RichContentDraftDto) error {
    r.rc = rc
    r.Set("rc", rc)
    return nil
}

func (r AlibabaAlihouseNewhomeRcSyncRequest) GetRc() *RichContentDraftDto {
    return r.rc
}

