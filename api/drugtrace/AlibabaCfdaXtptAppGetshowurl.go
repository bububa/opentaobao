package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
协同平台码查询页面url 
alibaba.cfda.xtpt.app.getshowurl

协同平台码查询页面url
*/
func AlibabaCfdaXtptAppGetshowurl(clt *core.SDKClient, req *drugtrace.AlibabaCfdaXtptAppGetshowurlRequest, session string) (*drugtrace.AlibabaCfdaXtptAppGetshowurlAPIResponse, error) {
    var resp drugtrace.AlibabaCfdaXtptAppGetshowurlAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
