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
    Response *TaobaoRdsDbGetdbResponse `json:"taobao_rds_db_getdb_response,omitempty"`
}

type TaobaoRdsDbGetdbResponse struct {

    // result
    Result   *TaobaoRdsDbGetdbResultSet `json:"result,omitempty"`

}
