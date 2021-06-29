package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询某个用户图片空间的使用情况 API请求
taobao.vas.service.getServTimes

查询某个用户图片空间的使用情况
*/
type TaobaoVasServiceGetServTimesRequest struct {
    model.Params
    // 服务编码
    _servCode   string
}

// 初始化TaobaoVasServiceGetServTimesRequest对象
func NewTaobaoVasServiceGetServTimesRequest() *TaobaoVasServiceGetServTimesRequest{
    return &TaobaoVasServiceGetServTimesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVasServiceGetServTimesRequest) GetApiMethodName() string {
    return "taobao.vas.service.getServTimes"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVasServiceGetServTimesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServCode Setter
// 服务编码
func (r *TaobaoVasServiceGetServTimesRequest) SetServCode(_servCode string) error {
    r._servCode = _servCode
    r.Set("serv_code", _servCode)
    return nil
}

// ServCode Getter
func (r TaobaoVasServiceGetServTimesRequest) GetServCode() string {
    return r._servCode
}
