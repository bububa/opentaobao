package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
货主仓库资源查询接口 API返回值 
taobao.qimen.warehouseinfo.query

货主仓库资源查询
*/
type TaobaoQimenWarehouseinfoQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenWarehouseinfoQueryAPIResponseModel
}

// 货主仓库资源查询接口 成功返回结果
type TaobaoQimenWarehouseinfoQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"qimen_warehouseinfo_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *TaobaoQimenWarehouseinfoQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}
