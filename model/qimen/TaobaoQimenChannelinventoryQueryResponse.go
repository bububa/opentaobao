package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
渠道库存查询接口 APIResponse
taobao.qimen.channelinventory.query

渠道库存查询
*/
type TaobaoQimenChannelinventoryQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenChannelinventoryQueryResponse `json:"taobao_qimen_channelinventory_query_response,omitempty"`
}

type TaobaoQimenChannelinventoryQueryResponse struct {

    // 
    Response   *ResponseDO `json:"response,omitempty"`

}
