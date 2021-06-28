package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
rds创建数据库 APIResponse
taobao.rds.db.create

在rds实例里创建数据库
*/
type TaobaoRdsDbCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRdsDbCreateResponse `json:"rds_db_create_response,omitempty"` 
    TaobaoRdsDbCreateResponse
}

/* model for simplify = false
type TaobaoRdsDbCreateResponse struct {

    // rds创建、查询、删除db返回结果的数据结构
    
    RdsDbInfo  *struct {
        RdsDbInfo  *RdsDbInfo `json:"rds_db_info,omitempty"`
    } `json:"rds_db_info,omitempty"`
    

}
*/

type TaobaoRdsDbCreateResponse struct {

    // rds创建、查询、删除db返回结果的数据结构
    RdsDbInfo   *RdsDbInfo `json:"rds_db_info,omitempty"`

}
