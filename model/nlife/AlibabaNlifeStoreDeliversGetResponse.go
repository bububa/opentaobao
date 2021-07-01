package nlife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取门店采购单下的发货单列表 API返回值 
alibaba.nlife.store.delivers.get

获取门店采购单下的发货单列表
*/
type AlibabaNlifeStoreDeliversGetAPIResponse struct {
    model.CommonResponse
    AlibabaNlifeStoreDeliversGetResponse
}

// 获取门店采购单下的发货单列表 成功返回结果
type AlibabaNlifeStoreDeliversGetResponse struct {
    XMLName xml.Name `xml:"alibaba_nlife_store_delivers_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询结果数据
    Data   *DeliverResponseDo `json:"data,omitempty" xml:"data,omitempty"`
}
