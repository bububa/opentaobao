package lstpos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV条码库查询接口 APIResponse
alibaba.lst.pos.open.goods.getgoodsbybarcode

ISV条码库查询接口
*/
type AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIResponse struct {
    model.CommonResponse
    AlibabaLstPosOpenGoodsGetgoodsbybarcodeResponse
}

type AlibabaLstPosOpenGoodsGetgoodsbybarcodeResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_pos_open_goods_getgoodsbybarcode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果对象
    
    Result   *AlibabaLstPosOpenGoodsGetgoodsbybarcodeResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
