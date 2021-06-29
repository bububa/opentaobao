package nrpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
去前置机pos商品离线文件下载地址查询接口 API请求
alibaba.mos.commdy.offline.getfileurl

去前置机-pos查询离线文件下载地址接口
*/
type AlibabaMosCommdyOfflineGetfileurlRequest struct {
    model.Params
    // 离线文件名称
    _fileKeys   []string
}

// 初始化AlibabaMosCommdyOfflineGetfileurlRequest对象
func NewAlibabaMosCommdyOfflineGetfileurlRequest() *AlibabaMosCommdyOfflineGetfileurlRequest{
    return &AlibabaMosCommdyOfflineGetfileurlRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosCommdyOfflineGetfileurlRequest) GetApiMethodName() string {
    return "alibaba.mos.commdy.offline.getfileurl"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosCommdyOfflineGetfileurlRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FileKeys Setter
// 离线文件名称
func (r *AlibabaMosCommdyOfflineGetfileurlRequest) SetFileKeys(_fileKeys []string) error {
    r._fileKeys = _fileKeys
    r.Set("file_keys", _fileKeys)
    return nil
}

// FileKeys Getter
func (r AlibabaMosCommdyOfflineGetfileurlRequest) GetFileKeys() []string {
    return r._fileKeys
}
