package nrpos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
去前置机商品在线查询 API返回值 
alibaba.mos.commdy.posmerchandise.getmerchandise

去前置机商品在线查询接口
*/
type AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse struct {
    model.CommonResponse
    AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponseModel
}

// 去前置机商品在线查询 成功返回结果
type AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mos_commdy_posmerchandise_getmerchandise_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *AlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
