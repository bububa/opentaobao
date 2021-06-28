package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
算法获取hscode APIResponse
tmall.item.calculate.hscode.get

算法获取hscode
*/
type TmallItemCalculateHscodeGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_calculate_hscode_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 算法返回预测的hscode数据
    
    Results   []string `json:"results,omitempty" xml:"