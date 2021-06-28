package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询骑手当前位置 APIResponse
alibaba.ele.fengniao.carrierdriver.location

查询骑手当前位置
*/
type AlibabaEleFengniaoCarrierdriverLocationAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ele_fengniao_carrierdriver_location_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // location
    
    Location   *Location `json:"location,omitempty" xml:"