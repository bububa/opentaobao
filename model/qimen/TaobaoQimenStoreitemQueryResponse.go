package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店关联商品查询接口 APIResponse
taobao.qimen.storeitem.query

商家在ERP等系统中调用该接口，查询某门店所关联的线上商品列表
*/
type TaobaoQimenStoreitemQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_storeitem_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 响应消息
    
    Message   string `json:"message,omitempty" xml:"