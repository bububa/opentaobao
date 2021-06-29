package film

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据手机查询匹配账号列表 API返回值 
taobao.film.account.phone.query

根据手机号查询匹配的账号列表
*/
type TaobaoFilmAccountPhoneQueryAPIResponse struct {
    model.CommonResponse
    TaobaoFilmAccountPhoneQueryResponse
}

// 根据手机查询匹配账号列表 成功返回结果
type TaobaoFilmAccountPhoneQueryResponse struct {
    XMLName xml.Name `xml:"film_account_phone_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回对象
    Result   *ResultListModel `json:"result,omitempty" xml:"result,omitempty"`
}
