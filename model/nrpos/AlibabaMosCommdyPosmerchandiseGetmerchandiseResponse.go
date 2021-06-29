package nrpos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
去前置机商品在线查询 APIResponse
alibaba.mos.commdy.posmerchandise.getmerchandise

去前置机商品在线查询接口
*/
type AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse struct {
    model.CommonResponse
    AlibabaMosCommdyPosmerchandiseGetmerchandiseResponse
}

type AlibabaMosCommdyPosmerchandiseGetmerchandiseResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_commdy_posmerchandise_getmerchandise_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
