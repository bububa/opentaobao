package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
rds创建数据库 APIResponse
taobao.rds.db.create

在rds实例里创建数据库
*/
type TaobaoRdsDbCreateAPIResponse struct {
    model.CommonResponse
    TaobaoRdsDbCreateResponse
}

type TaobaoRdsDbCreateResponse struct {
    XMLName xml.Name `xml:"rds_db_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // rds创建、查询、删除db返回结果的数据结构
    
    RdsDbInfo   *RdsDbInfo `json:"rds_db_info,omitempty" xml:"rds_db_info,omitempty"`

    
}
