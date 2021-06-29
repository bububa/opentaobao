package perfect

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同城购定制发品编辑 APIResponse
alibaba.perfect.performance.localitem.edit

同城购业务定制化发品接口，同城购业务线专用
*/
type AlibabaPerfectPerformanceLocalitemEditAPIResponse struct {
    model.CommonResponse
    AlibabaPerfectPerformanceLocalitemEditResponse
}

type AlibabaPerfectPerformanceLocalitemEditResponse struct {
    XMLName xml.Name `xml:"alibaba_perfect_performance_localitem_edit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回包装类
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
