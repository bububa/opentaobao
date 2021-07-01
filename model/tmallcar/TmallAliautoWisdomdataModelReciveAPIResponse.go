package tmallcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方车型数据上传 API返回值 
tmall.aliauto.wisdomdata.model.recive

天猫汽车对外提供的汽车车型数据上传接口
*/
type TmallAliautoWisdomdataModelReciveAPIResponse struct {
    model.CommonResponse
    TmallAliautoWisdomdataModelReciveAPIResponseModel
}

// 第三方车型数据上传 成功返回结果
type TmallAliautoWisdomdataModelReciveAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_aliauto_wisdomdata_model_recive_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TmallAliautoWisdomdataModelReciveResult `json:"result,omitempty" xml:"result,omitempty"`
}
