package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里房产图文草稿信息同步 API请求
alibaba.alihouse.newhome.rc.sync

接收图文草稿信息
*/
type AlibabaAlihouseNewhomeRcSyncRequest struct {
    model.Params
    // 图文内容
    _rc   *RichContentDraftDTO
}

// 初始化AlibabaAlihouseNewhomeRcSyncRequest对象
func NewAlibabaAlihouseNewhomeRcSyncRequest() *AlibabaAlihouseNewhomeRcSyncRequest{
    return &AlibabaAlihouseNewhomeRcSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeRcSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.rc.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeRcSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rc Setter
// 图文内容
func (r *AlibabaAlihouseNewhomeRcSyncRequest) SetRc(_rc *RichContentDraftDTO) error {
    r._rc = _rc
    r.Set("rc", _rc)
    return nil
}

// Rc Getter
func (r AlibabaAlihouseNewhomeRcSyncRequest) GetRc() *RichContentDraftDTO {
    return r._rc
}
