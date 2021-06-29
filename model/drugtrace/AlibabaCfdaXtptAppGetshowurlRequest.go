package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
协同平台码查询页面url APIRequest
alibaba.cfda.xtpt.app.getshowurl

协同平台码查询页面url
*/
type AlibabaCfdaXtptAppGetshowurlRequest struct {
    model.Params

    // 码
    code   string 

}

func NewAlibabaCfdaXtptAppGetshowurlRequest() *AlibabaCfdaXtptAppGetshowurlRequest{
    return &AlibabaCfdaXtptAppGetshowurlRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCfdaXtptAppGetshowurlRequest) GetApiMethodName() string {
    return "alibaba.cfda.xtpt.app.getshowurl"
}

func (r AlibabaCfdaXtptAppGetshowurlRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCfdaXtptAppGetshowurlRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaCfdaXtptAppGetshowurlRequest) GetCode() string {
    return r.code
}

