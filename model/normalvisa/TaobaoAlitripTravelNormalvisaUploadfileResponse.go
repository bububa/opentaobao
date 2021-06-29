package normalvisa

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传电子签证 API返回值 
taobao.alitrip.travel.normalvisa.uploadfile

上传电子签证
*/
type TaobaoAlitripTravelNormalvisaUploadfileAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelNormalvisaUploadfileResponse
}

// 上传电子签证 成功返回结果
type TaobaoAlitripTravelNormalvisaUploadfileResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_normalvisa_uploadfile_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果：包含results则代表成功保存
    Result   *TaobaoAlitripTravelNormalvisaUploadfileResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
