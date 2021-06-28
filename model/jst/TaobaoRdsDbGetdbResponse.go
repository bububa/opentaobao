package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
rds获取RDS的DB APIResponse
taobao.rds.db.getdb

rds获取RDS的DB
*/
type TaobaoRdsDbGetdbAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRdsDbGetdbResponse `json:"rds_db_getdb_response,omitempty"` 
    TaobaoRdsDbGetdbResponse
}

/* model for simplify = false
type TaobaoRdsDbGetdbResponse struct {

    // result
    
    Result  *struct {
        TaobaoRdsDbGetdbResultSet  *TaobaoRdsDbGetdbResultSet `json:"taobao_rds_db_getdb_result_set,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoRdsDbGetdbResponse struct {

    // result
    Result   *TaobaoRdsDbGetdbResultSet `json:"result,omitempty"`

}
