package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
Open Account索引查询 APIResponse
taobao.open.account.index.find

Open Account索引查询
*/
type TaobaoOpenAccountIndexFindAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenAccountIndexFindResponse `json:"taobao_open_account_index_find_response,omitempty"`
}

type TaobaoOpenAccountIndexFindResponse struct {

    // 返回结果
    Result   *OpenAccountResult `json:"result,omitempty"`

}
