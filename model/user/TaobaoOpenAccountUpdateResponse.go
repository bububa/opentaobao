package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Open Account数据更新 APIResponse
taobao.open.account.update

Open Account数据更新
*/
type TaobaoOpenAccountUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoOpenAccountUpdateResponse
}

type TaobaoOpenAccountUpdateResponse struct {
    XMLName xml.Name `xml:"open_account_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // update是否成功
    
    Datas   []OpenaccountVoid `json:"datas,omitempty" xml:"datas>openaccount_void,omitempty"`
    
    
}
