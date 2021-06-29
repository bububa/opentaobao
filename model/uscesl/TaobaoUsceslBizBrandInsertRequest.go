package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增电子价签商家 APIRequest
taobao.uscesl.biz.brand.insert

一个电子价签业务身份下新增商家接口
*/
type TaobaoUsceslBizBrandInsertRequest struct {
    model.Params

    // 商家名称
    brandName   string 

    // 商家外部编号
    brandOutCode   string 

}

func NewTaobaoUsceslBizBrandInsertRequest() *TaobaoUsceslBizBrandInsertRequest{
    return &TaobaoUsceslBizBrandInsertRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsceslBizBrandInsertRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.brand.insert"
}

func (r TaobaoUsceslBizBrandInsertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsceslBizBrandInsertRequest) SetBrandName(brandName string) error {
    r.brandName = brandName
    r.Set("brand_name", brandName)
    return nil
}

func (r TaobaoUsceslBizBrandInsertRequest) GetBrandName() string {
    return r.brandName
}

func (r *TaobaoUsceslBizBrandInsertRequest) SetBrandOutCode(brandOutCode string) error {
    r.brandOutCode = brandOutCode
    r.Set("brand_out_code", brandOutCode)
    return nil
}

func (r TaobaoUsceslBizBrandInsertRequest) GetBrandOutCode() string {
    return r.brandOutCode
}

