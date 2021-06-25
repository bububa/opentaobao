package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
货主仓库资源查询接口 APIResponse
taobao.qimen.warehouseinfo.query

货主仓库资源查询
*/
type TaobaoQimenWarehouseinfoQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenWarehouseinfoQueryResponse `json:"taobao_qimen_warehouseinfo_query_response,omitempty"`
}

type TaobaoQimenWarehouseinfoQueryResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
