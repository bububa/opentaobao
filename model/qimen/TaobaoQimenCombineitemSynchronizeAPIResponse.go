package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品接口 API返回值 
taobao.qimen.combineitem.synchronize

ERP调用奇门的接口,将商品信息同步给WMS
*/
type TaobaoQimenCombineitemSynchronizeAPIResponse struct {
    model.CommonResponse
    TaobaoQimenCombineitemSynchronizeAPIResponseModel
}

// 组合商品接口 成功返回结果
type TaobaoQimenCombineitemSynchronizeAPIResponseModel struct {
    XMLName xml.Name `xml:"qimen_combineitem_synchronize_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *TaobaoQimenCombineitemSynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}
