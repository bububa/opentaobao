package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品关联门店查询接口 APIResponse
taobao.qimen.itemstore.query

商家在ERP等系统中调用该接口，查询线上商品所关联的门店列表
*/
type TaobaoQimenItemstoreQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_itemstore_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 响应信息
    
    Message   string `json:"message,omitempty" xml:"