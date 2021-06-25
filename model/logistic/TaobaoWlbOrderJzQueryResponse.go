package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
家装业务查询物流公司api APIResponse
taobao.wlb.order.jz.query

家装业务查询物流公司api
*/
type TaobaoWlbOrderJzQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbOrderJzQueryResponse `json:"taobao_wlb_order_jz_query_response,omitempty"`
}

type TaobaoWlbOrderJzQueryResponse struct {

    // 错误编码
    ResultErrorCode   string `json:"result_error_code,omitempty"`

    // 错误信息
    ResultErrorMsg   string `json:"result_error_msg,omitempty"`

    // 结果信息
    Result   *JzTopDto `json:"result,omitempty"`

    // 是否成功
    ResultSuccess   bool `json:"result_success,omitempty"`

}
