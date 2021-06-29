package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
2-IBP查询CDC和RDC数据接口 APIRequest
alibaba.tmallgenie.scp.location.get

天猫精灵供应链-计划域-IBP查询CDC和RDC数据接口
*/
type AlibabaTmallgenieScpLocationGetRequest struct {
    model.Params

    // 扩展参数
    requestExtendJson   string 

}

func NewAlibabaTmallgenieScpLocationGetRequest() *AlibabaTmallgenieScpLocationGetRequest{
    return &AlibabaTmallgenieScpLocationGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpLocationGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.location.get"
}

func (r AlibabaTmallgenieScpLocationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpLocationGetRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

func (r AlibabaTmallgenieScpLocationGetRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}

