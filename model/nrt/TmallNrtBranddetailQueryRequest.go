package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌详情查询 APIRequest
tmall.nrt.branddetail.query

根据品牌id查询品牌的详细信息
*/
type TmallNrtBranddetailQueryRequest struct {
    model.Params

    // 品牌id
    brandId   int64 

}

func NewTmallNrtBranddetailQueryRequest() *TmallNrtBranddetailQueryRequest{
    return &TmallNrtBranddetailQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrtBranddetailQueryRequest) GetApiMethodName() string {
    return "tmall.nrt.branddetail.query"
}

func (r TmallNrtBranddetailQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrtBranddetailQueryRequest) SetBrandId(brandId int64) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

func (r TmallNrtBranddetailQueryRequest) GetBrandId() int64 {
    return r.brandId
}

