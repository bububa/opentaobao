package baodian

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取宝点SDK的配置项（已迁移） APIResponse
taobao.baodian.server.sdk.config.get

获取SDK各种配置项（已迁移）
*/
type TaobaoBaodianServerSdkConfigGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaodianServerSdkConfigGetResponse `json:"taobao_baodian_server_sdk_config_get_response,omitempty"`
}

type TaobaoBaodianServerSdkConfigGetResponse struct {

    // 返回sdk配置的字符串
    Result   string `json:"result,omitempty"`

}
