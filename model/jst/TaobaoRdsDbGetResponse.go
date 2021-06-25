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
    Response *TaobaoRdsDbGetResponse `json:"taobao_rds_db_get_response,omitempty"`
}

type TaobaoRdsDbGetResponse struct {

    // 数据库查询返回数据结构
    RdsDbInfos   []RdsDbInfo `json:"rds_db_infos,omitempty"`

}
