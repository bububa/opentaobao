package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
rds获取RDS的DB API返回值 
taobao.rds.db.getdb

rds获取RDS的DB
*/
type TaobaoRdsDbGetdbAPIResponse struct {
    model.CommonResponse
    TaobaoRdsDbGetdbResponse
}

// rds获取RDS的DB 成功返回结果
type TaobaoRdsDbGetdbResponse struct {
    XMLName xml.Name `xml:"rds_db_getdb_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoRdsDbGetdbResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
