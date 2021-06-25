package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出库单确认接口 APIRequest
taobao.qimen.stockout.confirm

货品出库后，WMS将状态回传给ERP
*/
type TaobaoQimenStockoutConfirmRequest struct {
    model.Params

    // 
    request   *TaobaoQimenStockoutConfirmStruct 

}

func NewTaobaoQimenStockoutConfirmRequest() *TaobaoQimenStockoutConfirmRequest{
    return &TaobaoQimenStockoutConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenStockoutConfirmRequest) GetApiMethodName() string {
    return "taobao.qimen.stockout.confirm"
}

func (r TaobaoQimenStockoutConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenStockoutConfirmRequest) SetRequest(request *TaobaoQimenStockoutConfirmStruct) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenStockoutConfirmRequest) GetRequest() *TaobaoQimenStockoutConfirmStruct {
    return r.request
}

