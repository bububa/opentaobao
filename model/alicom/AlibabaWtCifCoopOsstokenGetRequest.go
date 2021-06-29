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
    _appName   string
    // 系统分配的source
    _source   string
    // 系统分配的biz
    _biz   string
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
func (r *AlibabaWtCifCoopOsstokenGetRequest) SetAppName(_appName string) error {
    r._appName = _appName
    r.Set("app_name", _appName)
    return nil
}

// AppName Getter
func (r AlibabaWtCifCoopOsstokenGetRequest) GetAppName() string {
    return r._appName
}
// Source Setter
// 系统分配的source
func (r *AlibabaWtCifCoopOsstokenGetRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r AlibabaWtCifCoopOsstokenGetRequest) GetSource() string {
    return r._source
}
// Biz Setter
// 系统分配的biz
func (r *AlibabaWtCifCoopOsstokenGetRequest) SetBiz(_biz string) error {
    r._biz = _biz
    r.Set("biz", _biz)
    return nil
}

// Biz Getter
func (r AlibabaWtCifCoopOsstokenGetRequest) GetBiz() string {
    return r._biz
}
