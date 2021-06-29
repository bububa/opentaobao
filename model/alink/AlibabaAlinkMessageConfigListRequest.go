package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询消息开关配置列表 API请求
alibaba.alink.message.config.list

阿里智能获取消息开关配置列表
*/
type AlibabaAlinkMessageConfigListRequest struct {
    model.Params
}

// 初始化AlibabaAlinkMessageConfigListRequest对象
func NewAlibabaAlinkMessageConfigListRequest() *AlibabaAlinkMessageConfigListRequest{
    return &AlibabaAlinkMessageConfigListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlinkMessageConfigListRequest) GetApiMethodName() string {
    return "alibaba.alink.message.config.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlinkMessageConfigListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
