package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取oss签名接口 API请求
alibaba.wt.cif.coop.osstoken.get

商家合作上传oss图片获取token接口
*/
type AlibabaWtCifCoopOsstokenGetRequest struct {
    model.Params
    // 调用方的应用名
    appName   string
    // 系统分配的source
    source   string
    // 系统分配的biz
    biz   string
}

// 初始化AlibabaWtCifCoopOsstokenGetRequest对象
func NewAlibabaWtCifCoopOsstokenGetRequest() *AlibabaWtCifCoopOsstokenGetRequest{
    return &AlibabaWtCifCoopOsstokenGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWtCifCoopOsstokenGetRequest) GetApiMethodName() string {
    return "alibaba.wt.cif.coop.osstoken.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWtCifCoopOsstokenGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppName Setter
// 调用方的应用名
func (r *AlibabaWtCifCoopOsstokenGetRequest) SetAppName(appName string) error {
    r.appName = appName
    r.Set("app_name", appName)
    return nil
}

// AppName Getter
func (r AlibabaWtCifCoopOsstokenGetRequest) GetAppName() string {
    return r.appName
}
// Source Setter
// 系统分配的source
func (r *AlibabaWtCifCoopOsstokenGetRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r AlibabaWtCifCoopOsstokenGetRequest) GetSource() string {
    return r.source
}
// Biz Setter
// 系统分配的biz
func (r *AlibabaWtCifCoopOsstokenGetRequest) SetBiz(biz string) error {
    r.biz = biz
    r.Set("biz", biz)
    return nil
}

// Biz Getter
func (r AlibabaWtCifCoopOsstokenGetRequest) GetBiz() string {
    return r.biz
}
