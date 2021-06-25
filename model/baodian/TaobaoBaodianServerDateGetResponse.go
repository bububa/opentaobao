package baodian

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务器时间获取 APIResponse
taobao.baodian.server.date.get

获取服务器时间
*/
type TaobaoBaodianServerDateGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaodianServerDateGetResponse `json:"taobao_baodian_server_date_get_response,omitempty"`
}

type TaobaoBaodianServerDateGetResponse struct {

    // 返回时间为毫秒
    ServerDate   int64 `json:"server_date,omitempty"`

}
