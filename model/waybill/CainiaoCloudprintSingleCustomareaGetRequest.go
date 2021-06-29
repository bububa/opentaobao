package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家单一自定义区 API请求
cainiao.cloudprint.single.customarea.get

商家所有快递公司模板只有一个自定义区
*/
type CainiaoCloudprintSingleCustomareaGetRequest struct {
    model.Params
    // 这是商家用户id
    _sellerId   int64
}

// 初始化CainiaoCloudprintSingleCustomareaGetRequest对象
func NewCainiaoCloudprintSingleCustomareaGetRequest() *CainiaoCloudprintSingleCustomareaGetRequest{
    return &CainiaoCloudprintSingleCustomareaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintSingleCustomareaGetRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.single.customarea.get"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintSingleCustomareaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SellerId Setter
// 这是商家用户id
func (r *CainiaoCloudprintSingleCustomareaGetRequest) SetSellerId(_sellerId int64) error {
    r._sellerId = _sellerId
    r.Set("seller_id", _sellerId)
    return nil
}

// SellerId Getter
func (r CainiaoCloudprintSingleCustomareaGetRequest) GetSellerId() int64 {
    return r._sellerId
}
