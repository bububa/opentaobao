package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单个删除订阅关系 APIResponse
taobao.baichuan.item.unsubscribe

删除单个商品订阅关系
*/
type TaobaoBaichuanItemUnsubscribeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"baichuan_item_unsubscribe_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoBaichuanItemUnsubscribeResult `json:"result,omitempty" xml:"