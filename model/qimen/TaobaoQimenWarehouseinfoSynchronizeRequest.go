package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓库同步接口 APIRequest
taobao.qimen.warehouseinfo.synchronize

仓库同步接口
*/
type TaobaoQimenWarehouseinfoSynchronizeRequest struct {
    model.Params

    // 请求报文
    request   *Request 

}

func NewTaobaoQimenWarehouseinfoSynchronizeRequest() *TaobaoQimenWarehouseinfoSynchronizeRequest{
    return &TaobaoQimenWarehouseinfoSynchronizeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenWarehouseinfoSynchronizeRequest) GetApiMethodName() string {
    return "taobao.qimen.warehouseinfo.synchronize"
}

func (r TaobaoQimenWarehouseinfoSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenWarehouseinfoSynchronizeRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenWarehouseinfoSynchronizeRequest) GetRequest() *Request {
    return r.request
}

