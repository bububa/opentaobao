package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取oss签名接口 APIRequest
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

func NewAlibabaWtCifCoopOsstokenGetRequest() *AlibabaWtCifCoopOsstokenGetRequest{
    return &AlibabaWtCifCoopOsstokenGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWtCifCoopOsstokenGetRequest) GetApiMethodName() string {
    return "alibaba.wt.cif.coop.osstoken.get"
}

func (r AlibabaWtCifCoopOsstokenGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWtCifCoopOsstokenGetRequest) SetAppName(appName string) error {
    r.appName = appName
    r.Set("app_name", appName)
    return nil
}

func (r AlibabaWtCifCoopOsstokenGetRequest) GetAppName() string {
    return r.appName
}

func (r *AlibabaWtCifCoopOsstokenGetRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r AlibabaWtCifCoopOsstokenGetRequest) GetSource() string {
    return r.source
}

func (r *AlibabaWtCifCoopOsstokenGetRequest) SetBiz(biz string) error {
    r.biz = biz
    r.Set("biz", biz)
    return nil
}

func (r AlibabaWtCifCoopOsstokenGetRequest) GetBiz() string {
    return r.biz
}

