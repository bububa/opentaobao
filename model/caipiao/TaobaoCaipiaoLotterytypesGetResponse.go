package caipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取可用的彩种列表 APIResponse
taobao.caipiao.lotterytypes.get

获取彩票系统支持的可用于赠送的彩种列表
*/
type TaobaoCaipiaoLotterytypesGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"caipiao_lotterytypes_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 彩种个数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"