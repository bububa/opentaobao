package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取TOP通道解密秘钥 API返回值 
taobao.top.secret.get

top sdk通过api获取对应解密秘钥
*/
type TaobaoTopSecretGetAPIResponse struct {
    model.CommonResponse
    TaobaoTopSecretGetResponse
}

// 获取TOP通道解密秘钥 成功返回结果
type TaobaoTopSecretGetResponse struct {
    XMLName xml.Name `xml:"top_secret_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 下次更新秘钥间隔，单位（秒）
    Interval   int64 `json:"interval,omitempty" xml:"interval,omitempty"`
    // 秘钥值
    Secret   string `json:"secret,omitempty" xml:"secret,omitempty"`
    // 秘钥版本号
    SecretVersion   int64 `json:"secret_version,omitempty" xml:"secret_version,omitempty"`
    // 最长有效期，容灾使用，单位（秒）
    MaxInterval   int64 `json:"max_interval,omitempty" xml:"max_interval,omitempty"`
    // app配置信息
    AppConfig   string `json:"app_config,omitempty" xml:"app_config,omitempty"`
}
