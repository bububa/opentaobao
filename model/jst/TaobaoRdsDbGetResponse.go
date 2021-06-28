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
    TaobaoRdsDbGetResponse
}

type TaobaoRdsDbGetResponse struct {
    XMLName xml.Name `xml:"rds_db_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 数据库查询返回数据结构
    
    RdsDbInfos   []RdsDbInfo `json:"rds_db_infos,omitempty" xml:"rds_db_infos>rds_db_info,omitempty"`
    
    
}
