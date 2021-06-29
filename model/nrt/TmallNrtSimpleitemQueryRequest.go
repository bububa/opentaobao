package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
简易商品查询接口 APIRequest
tmall.nrt.simpleitem.query

为居然之家和阿里的合资公司 homeStyler提供简易的商品信息查询 包含商品名称  图片 状态

后续合资公司服务会迁移到内网 暂时过渡用
*/
type TmallNrtSimpleitemQueryRequest struct {
    model.Params

    // 商品编码数组
    ids   []int64 

}

func NewTmallNrtSimpleitemQueryRequest() *TmallNrtSimpleitemQueryRequest{
    return &TmallNrtSimpleitemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrtSimpleitemQueryRequest) GetApiMethodName() string {
    return "tmall.nrt.simpleitem.query"
}

func (r TmallNrtSimpleitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrtSimpleitemQueryRequest) SetIds(ids []int64) error {
    r.ids = ids
    r.Set("ids", ids)
    return nil
}

func (r TmallNrtSimpleitemQueryRequest) GetIds() []int64 {
    return r.ids
}

