package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获得创意 APIResponse
taobao.simba.creatives.get

取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；<br/>如果同时提供了推广组Id和创意id列表，则优先使用推广组Id；
*/
type TaobaoSimbaCreativesGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_creatives_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 创意对象列表
    
    Creatives   []Creative `json:"creatives,omitempty" xml:"