package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询rds下的数据库 APIResponse
taobao.rds.db.get

查询rds实例下的数据库
*/
type TaobaoRdsDbGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"rds_db_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 数据库查询返回数据结构
    
    RdsDbInfos   []RdsDbInfo `json:"rds_db_infos,omitempty" xml:"