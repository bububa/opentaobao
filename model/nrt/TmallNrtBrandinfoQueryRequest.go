package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌数据查询 APIRequest
tmall.nrt.brandinfo.query

商家获取自己旗舰店授权的品牌id列表
*/
type TmallNrtBrandinfoQueryRequest struct {
    model.Params

}

func NewTmallNrtBrandinfoQueryRequest() *TmallNrtBrandinfoQueryRequest{
    return &TmallNrtBrandinfoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrtBrandinfoQueryRequest) GetApiMethodName() string {
    return "tmall.nrt.brandinfo.query"
}

func (r TmallNrtBrandinfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


