package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品删除 APIResponse
alibaba.item.operate.delete

商品删除
*/
type AlibabaItemOperateDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_item_operate_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品删除是否成功
    
    Result   string `json:"result,omitempty" xml:"