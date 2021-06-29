package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
协同平台码查询页面url API请求
alibaba.cfda.xtpt.app.getshowurl

协同平台码查询页面url
*/
type AlibabaCfdaXtptAppGetshowurlRequest struct {
    model.Params
    // 码
    code   string
}

// 初始化AlibabaCfdaXtptAppGetshowurlRequest对象
func NewAlibabaCfdaXtptAppGetshowurlRequest() *AlibabaCfdaXtptAppGetshowurlRequest{
    return &AlibabaCfdaXtptAppGetshowurlRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCfdaXtptAppGetshowurlRequest) GetApiMethodName() string {
    return "alibaba.cfda.xtpt.app.getshowurl"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCfdaXtptAppGetshowurlRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 码
func (r *AlibabaCfdaXtptAppGetshowurlRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaCfdaXtptAppGetshowurlRequest) GetCode() string {
    return r.code
}
