package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
RDS数据库删除 APIResponse
taobao.rds.db.delete

通过api删除用户RDS的数据库
*/
type TaobaoRdsDbDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoRdsDbDeleteResponse `json:"taobao_rds_db_delete_response,omitempty"`
}

type TaobaoRdsDbDeleteResponse struct {

    // 删除数据库，返回结果对象
    RdsDbInfo   *RdsDbInfo `json:"rds_db_info,omitempty"`

}
