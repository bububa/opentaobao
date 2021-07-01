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
type TmallNrtBrandinfoQueryAPIRequest struct {
    model.Params
}

// 初始化TmallNrtBrandinfoQueryAPIRequest对象
func NewTmallNrtBrandinfoQueryRequest() *TmallNrtBrandinfoQueryAPIRequest{
    return &TmallNrtBrandinfoQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtBrandinfoQueryAPIRequest) GetApiMethodName() string {
    return "tmall.nrt.brandinfo.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtBrandinfoQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
