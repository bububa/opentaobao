package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
rds创建数据库账户 APIResponse
taobao.rds.db.createaccount

rds创建数据库账户
*/
type TaobaoRdsDbCreateaccountAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRdsDbCreateaccountResponse `json:"rds_db_createaccount_response,omitempty"` 
    TaobaoRdsDbCreateaccountResponse
}

/* model for simplify = false
type TaobaoRdsDbCreateaccountResponse struct {

    // result
    
    Result  *struct {
        TaobaoRdsDbCreateaccountResultSet  *TaobaoRdsDbCreateaccountResultSet `json:"taobao_rds_db_createaccount_result_set,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoRdsDbCreateaccountResponse struct {

    // result
    Result   *TaobaoRdsDbCreateaccountResultSet `json:"result,omitempty"`

}
