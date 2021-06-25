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
    Response *TaobaoRdsDbCreateResponse `json:"taobao_rds_db_create_response,omitempty"`
}

type TaobaoRdsDbCreateResponse struct {

    // rds创建、查询、删除db返回结果的数据结构
    RdsDbInfo   *RdsDbInfo `json:"rds_db_info,omitempty"`

}
