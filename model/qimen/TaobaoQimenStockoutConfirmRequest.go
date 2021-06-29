package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出库单确认接口 API请求
taobao.qimen.stockout.confirm

货品出库后，WMS将状态回传给ERP
*/
type TaobaoQimenStockoutConfirmRequest struct {
    model.Params
    // 
    request   *TaobaoQimenStockoutConfirmStruct
}

// 初始化TaobaoQimenStockoutConfirmRequest对象
func NewTaobaoQimenStockoutConfirmRequest() *TaobaoQimenStockoutConfirmRequest{
    return &TaobaoQimenStockoutConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenStockoutConfirmRequest) GetApiMethodName() string {
    return "taobao.qimen.stockout.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenStockoutConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenStockoutConfirmRequest) SetRequest(request *TaobaoQimenStockoutConfirmStruct) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenStockoutConfirmRequest) GetRequest() *TaobaoQimenStockoutConfirmStruct {
    return r.request
}
