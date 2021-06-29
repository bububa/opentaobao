package tmallcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方车型数据上传 APIResponse
tmall.aliauto.wisdomdata.model.recive

天猫汽车对外提供的汽车车型数据上传接口
*/
type TmallAliautoWisdomdataModelReciveAPIResponse struct {
    model.CommonResponse
    TmallAliautoWisdomdataModelReciveResponse
}

type TmallAliautoWisdomdataModelReciveResponse struct {
    XMLName xml.Name `xml:"tmall_aliauto_wisdomdata_model_recive_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TmallAliautoWisdomdataModelReciveResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
