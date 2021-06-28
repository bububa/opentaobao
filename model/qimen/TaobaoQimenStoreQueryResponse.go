package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店信息查询接口 APIResponse
taobao.qimen.store.query

商家在ERP等系统中调用该接口，查询门店相关信息
*/
type TaobaoQimenStoreQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_store_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 门店名称
    
    StoreName   string `json:"store_name,omitempty" xml:"