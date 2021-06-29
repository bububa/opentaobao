package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据产品规格的Id号获取当个的规格信息 API请求
tmall.product.spec.get

通过当个的spec_id获取到这个产品规格的信息，主要是因为产品规格是要经过审核的，所以通过这个接口可以获取到是否通过审核<br/>通过参看这个ProductSpec的status判断：<br/>1:表示审核通过<br/>3:表示等待审核。<br/>如果你的id找不到数据，那么就是审核被拒绝。
*/
type TmallProductSpecGetRequest struct {
    model.Params
    // 要获取信息的产品规格信息。
    specId   int64
}

// 初始化TmallProductSpecGetRequest对象
func NewTmallProductSpecGetRequest() *TmallProductSpecGetRequest{
    return &TmallProductSpecGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallProductSpecGetRequest) GetApiMethodName() string {
    return "tmall.product.spec.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallProductSpecGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SpecId Setter
// 要获取信息的产品规格信息。
func (r *TmallProductSpecGetRequest) SetSpecId(specId int64) error {
    r.specId = specId
    r.Set("spec_id", specId)
    return nil
}

// SpecId Getter
func (r TmallProductSpecGetRequest) GetSpecId() int64 {
    return r.specId
}
