package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
rds获取RDS的DB APIResponse
taobao.rds.db.getdb

rds获取RDS的DB
*/
type TaobaoRdsDbGetdbAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"rds_db_getdb_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoRdsDbGetdbResultSet `json:"result,omitempty" xml:"