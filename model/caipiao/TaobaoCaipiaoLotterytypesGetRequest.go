package caipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取可用的彩种列表 APIRequest
taobao.caipiao.lotterytypes.get

获取彩票系统支持的可用于赠送的彩种列表
*/
type TaobaoCaipiaoLotterytypesGetRequest struct {
    model.Params

}

func NewTaobaoCaipiaoLotterytypesGetRequest() *TaobaoCaipiaoLotterytypesGetRequest{
    return &TaobaoCaipiaoLotterytypesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCaipiaoLotterytypesGetRequest) GetApiMethodName() string {
    return "taobao.caipiao.lotterytypes.get"
}

func (r TaobaoCaipiaoLotterytypesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


