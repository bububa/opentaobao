package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询rds下的数据库 APIResponse
taobao.rds.db.get

查询rds实例下的数据库
*/
type TaobaoRdsDbGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRdsDbGetResponse `json:"rds_db_get_response,omitempty"` 
    TaobaoRdsDbGetResponse
}

/* model for simplify = false
type TaobaoRdsDbGetResponse struct {

    // 数据库查询返回数据结构
    
    RdsDbInfos  struct {
        RdsDbInfo  []RdsDbInfo `json:"rds_db_info,omitempty"`
    } `json:"rds_db_infos,omitempty"`
    

}
*/

type TaobaoRdsDbGetResponse struct {

    // 数据库查询返回数据结构
    RdsDbInfos   []RdsDbInfo `json:"rds_db_infos,omitempty"`

}
