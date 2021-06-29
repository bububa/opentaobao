package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增电子价签商家 API请求
taobao.uscesl.biz.brand.insert

一个电子价签业务身份下新增商家接口
*/
type TaobaoUsceslBizBrandInsertRequest struct {
    model.Params
    // 商家名称
    _brandName   string
    // 商家外部编号
    _brandOutCode   string
}

// 初始化TaobaoUsceslBizBrandInsertRequest对象
func NewTaobaoUsceslBizBrandInsertRequest() *TaobaoUsceslBizBrandInsertRequest{
    return &TaobaoUsceslBizBrandInsertRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizBrandInsertRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.brand.insert"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizBrandInsertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BrandName Setter
// 商家名称
func (r *TaobaoUsceslBizBrandInsertRequest) SetBrandName(_brandName string) error {
    r._brandName = _brandName
    r.Set("brand_name", _brandName)
    return nil
}

// BrandName Getter
func (r TaobaoUsceslBizBrandInsertRequest) GetBrandName() string {
    return r._brandName
}
// BrandOutCode Setter
// 商家外部编号
func (r *TaobaoUsceslBizBrandInsertRequest) SetBrandOutCode(_brandOutCode string) error {
    r._brandOutCode = _brandOutCode
    r.Set("brand_out_code", _brandOutCode)
    return nil
}

// BrandOutCode Getter
func (r TaobaoUsceslBizBrandInsertRequest) GetBrandOutCode() string {
    return r._brandOutCode
}
