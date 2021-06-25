package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询某个用户图片空间的使用情况 APIRequest
taobao.vas.service.getServTimes

查询某个用户图片空间的使用情况
*/
type TaobaoVasServiceGetServTimesRequest struct {
    model.Params

    // 服务编码
    servCode   string 

}

func NewTaobaoVasServiceGetServTimesRequest() *TaobaoVasServiceGetServTimesRequest{
    return &TaobaoVasServiceGetServTimesRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVasServiceGetServTimesRequest) GetApiMethodName() string {
    return "taobao.vas.service.getServTimes"
}

func (r TaobaoVasServiceGetServTimesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVasServiceGetServTimesRequest) SetServCode(servCode string) error {
    r.servCode = servCode
    r.Set("serv_code", servCode)
    return nil
}

func (r TaobaoVasServiceGetServTimesRequest) GetServCode() string {
    return r.servCode
}

