package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌详情查询 API请求
tmall.nrt.branddetail.query

根据品牌id查询品牌的详细信息
*/
type TmallNrtBranddetailQueryRequest struct {
    model.Params
    // 品牌id
    _brandId   int64
}

// 初始化TmallNrtBranddetailQueryRequest对象
func NewTmallNrtBranddetailQueryRequest() *TmallNrtBranddetailQueryRequest{
    return &TmallNrtBranddetailQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtBranddetailQueryRequest) GetApiMethodName() string {
    return "tmall.nrt.branddetail.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtBranddetailQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BrandId Setter
// 品牌id
func (r *TmallNrtBranddetailQueryRequest) SetBrandId(_brandId int64) error {
    r._brandId = _brandId
    r.Set("brand_id", _brandId)
    return nil
}

// BrandId Getter
func (r TmallNrtBranddetailQueryRequest) GetBrandId() int64 {
    return r._brandId
}
