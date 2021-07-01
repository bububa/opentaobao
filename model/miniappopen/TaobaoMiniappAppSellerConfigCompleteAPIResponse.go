package miniappopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家完成小程序相关配置 API返回值 
taobao.miniapp.app.seller.config.complete

通过该接口告知平台商家已经完成小程序相关的必要设置，可进行后续操作。主要用于小部件、客服插件等场景。
*/
type TaobaoMiniappAppSellerConfigCompleteAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappAppSellerConfigCompleteAPIResponseModel
}

// 商家完成小程序相关配置 成功返回结果
type TaobaoMiniappAppSellerConfigCompleteAPIResponseModel struct {
    XMLName xml.Name `xml:"miniapp_app_seller_config_complete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 操作成功与否
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
}
