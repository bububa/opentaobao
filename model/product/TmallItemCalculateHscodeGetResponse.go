package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
算法获取hscode APIResponse
tmall.item.calculate.hscode.get

算法获取hscode
*/
type TmallItemCalculateHscodeGetAPIResponse struct {
    model.CommonResponse
    Response *TmallItemCalculateHscodeGetResponse `json:"tmall_item_calculate_hscode_get_response,omitempty"`
}

type TmallItemCalculateHscodeGetResponse struct {

    // 算法返回预测的hscode数据
    Results   []Json `json:"results,omitempty"`

}
