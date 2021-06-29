package fivee

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
进口商品查询 API返回值 
taobao.fivee.importproduct.get

资质共享平台查询进口商品信息
*/
type TaobaoFiveeImportproductGetAPIResponse struct {
    model.CommonResponse
    TaobaoFiveeImportproductGetResponse
}

// 进口商品查询 成功返回结果
type TaobaoFiveeImportproductGetResponse struct {
    XMLName xml.Name `xml:"fivee_importproduct_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoFiveeImportproductGetResult `json:"result,omitempty" xml:"result,omitempty"`
}