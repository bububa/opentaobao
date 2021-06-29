package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
2-IBP查询CDC和RDC数据接口 API请求
alibaba.tmallgenie.scp.location.get

天猫精灵供应链-计划域-IBP查询CDC和RDC数据接口
*/
type AlibabaTmallgenieScpLocationGetRequest struct {
    model.Params
    // 扩展参数
    requestExtendJson   string
}

// 初始化AlibabaTmallgenieScpLocationGetRequest对象
func NewAlibabaTmallgenieScpLocationGetRequest() *AlibabaTmallgenieScpLocationGetRequest{
    return &AlibabaTmallgenieScpLocationGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpLocationGetRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.location.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpLocationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpLocationGetRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

// RequestExtendJson Getter
func (r AlibabaTmallgenieScpLocationGetRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}
