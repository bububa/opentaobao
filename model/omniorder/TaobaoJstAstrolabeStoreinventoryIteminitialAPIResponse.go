package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
库存初始化接口 API返回值 
taobao.jst.astrolabe.storeinventory.iteminitial

ERP调用奇门的接口，对门店的库存进行初始化
*/
type TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse struct {
    model.CommonResponse
    TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponseModel
}

// 库存初始化接口 成功返回结果
type TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponseModel struct {
    XMLName xml.Name `xml:"jst_astrolabe_storeinventory_iteminitial_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 响应标示
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`
    // 响应标签
    QimenCode   string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`
    // 响应信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 错误信息列表
    ErrorDescriptions   []TaobaoJstAstrolabeStoreinventoryIteminitialError `json:"error_descriptions,omitempty" xml:"error_descriptions>taobao_jst_astrolabe_storeinventory_iteminitial_error,omitempty"`
}
