package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
RDS数据库删除 APIResponse
taobao.rds.db.delete

通过api删除用户RDS的数据库
*/
type TaobaoRdsDbDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"rds_db_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 删除数据库，返回结果对象
    
    RdsDbInfo   *RdsDbInfo `json:"rds_db_info,omitempty" xml:"