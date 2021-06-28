package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
rds创建数据库账户 APIResponse
taobao.rds.db.createaccount

rds创建数据库账户
*/
type TaobaoRdsDbCreateaccountAPIResponse struct {
    model.CommonResponse
    TaobaoRdsDbCreateaccountResponse
}

type TaobaoRdsDbCreateaccountResponse struct {
    XMLName xml.Name `xml:"rds_db_createaccount_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoRdsDbCreateaccountResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
