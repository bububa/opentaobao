package nrpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
去前置机pos商品离线文件下载地址查询接口 APIRequest
alibaba.mos.commdy.offline.getfileurl

去前置机-pos查询离线文件下载地址接口
*/
type AlibabaMosCommdyOfflineGetfileurlRequest struct {
    model.Params

    // 离线文件名称
    fileKeys   []string 

}

func NewAlibabaMosCommdyOfflineGetfileurlRequest() *AlibabaMosCommdyOfflineGetfileurlRequest{
    return &AlibabaMosCommdyOfflineGetfileurlRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosCommdyOfflineGetfileurlRequest) GetApiMethodName() string {
    return "alibaba.mos.commdy.offline.getfileurl"
}

func (r AlibabaMosCommdyOfflineGetfileurlRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosCommdyOfflineGetfileurlRequest) SetFileKeys(fileKeys []string) error {
    r.fileKeys = fileKeys
    r.Set("file_keys", fileKeys)
    return nil
}

func (r AlibabaMosCommdyOfflineGetfileurlRequest) GetFileKeys() []string {
    return r.fileKeys
}

