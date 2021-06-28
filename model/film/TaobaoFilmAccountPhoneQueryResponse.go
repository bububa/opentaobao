package film

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据手机查询匹配账号列表 APIResponse
taobao.film.account.phone.query

根据手机号查询匹配的账号列表
*/
type TaobaoFilmAccountPhoneQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"film_account_phone_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回对象
    
    Result   *ResultListModel `json:"result,omitempty" xml:"