package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌数据查询 API请求
tmall.nrt.brandinfo.query

商家获取自己旗舰店授权的品牌id列表
*/
type TmallNrtBrandinfoQueryRequest struct {
    model.Params
}

// 初始化TmallNrtBrandinfoQueryRequest对象
func NewTmallNrtBrandinfoQueryRequest() *TmallNrtBrandinfoQueryRequest{
    return &TmallNrtBrandinfoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtBrandinfoQueryRequest) GetApiMethodName() string {
    return "tmall.nrt.brandinfo.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtBrandinfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
