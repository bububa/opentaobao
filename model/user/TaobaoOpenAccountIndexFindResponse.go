package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Open Account索引查询 APIResponse
taobao.open.account.index.find

Open Account索引查询
*/
type TaobaoOpenAccountIndexFindAPIResponse struct {
    model.CommonResponse
    TaobaoOpenAccountIndexFindResponse
}

type TaobaoOpenAccountIndexFindResponse struct {
    XMLName xml.Name `xml:"open_account_index_find_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *OpenAccountResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
