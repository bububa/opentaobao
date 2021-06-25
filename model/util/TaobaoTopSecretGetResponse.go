package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取TOP通道解密秘钥 APIResponse
taobao.top.secret.get

top sdk通过api获取对应解密秘钥
*/
type TaobaoTopSecretGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTopSecretGetResponse `json:"taobao_top_secret_get_response,omitempty"`
}

type TaobaoTopSecretGetResponse struct {

    // 下次更新秘钥间隔，单位（秒）
    Interval   int64 `json:"interval,omitempty"`

    // 秘钥值
    Secret   string `json:"secret,omitempty"`

    // 秘钥版本号
    SecretVersion   int64 `json:"secret_version,omitempty"`

    // 最长有效期，容灾使用，单位（秒）
    MaxInterval   int64 `json:"max_interval,omitempty"`

    // app配置信息
    AppConfig   string `json:"app_config,omitempty"`

}
