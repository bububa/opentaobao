package alink

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alink"
)

/* 
开放数据授权访问URL查询 
aliyun.alink.opendata.url.query

厂商数据授权访问URL查询
*/
func AliyunAlinkOpendataUrlQuery(clt *core.SDKClient, req *alink.AliyunAlinkOpendataUrlQueryRequest, session string) (*alink.AliyunAlinkOpendataUrlQueryAPIResponse, error) {
    var resp alink.AliyunAlinkOpendataUrlQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
