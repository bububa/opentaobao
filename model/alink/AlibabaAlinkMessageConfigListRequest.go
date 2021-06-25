package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询消息开关配置列表 APIRequest
alibaba.alink.message.config.list

阿里智能获取消息开关配置列表
*/
type AlibabaAlinkMessageConfigListRequest struct {
    model.Params

}

func NewAlibabaAlinkMessageConfigListRequest() *AlibabaAlinkMessageConfigListRequest{
    return &AlibabaAlinkMessageConfigListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlinkMessageConfigListRequest) GetApiMethodName() string {
    return "alibaba.alink.message.config.list"
}

func (r AlibabaAlinkMessageConfigListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


