package lstpos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV条码库查询接口 API返回值 
alibaba.lst.pos.open.goods.getgoodsbybarcode

ISV条码库查询接口
*/
type AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIResponse struct {
    model.CommonResponse
    AlibabaLstPosOpenGoodsGetgoodsbybarcodeResponse
}

// ISV条码库查询接口 成功返回结果
type AlibabaLstPosOpenGoodsGetgoodsbybarcodeResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_pos_open_goods_getgoodsbybarcode_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果对象
    Result   *AlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
