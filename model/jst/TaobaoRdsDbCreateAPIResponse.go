package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdsDbCreateAPIResponse rds创建数据库 API返回值
// taobao.rds.db.create
//
// 在rds实例里创建数据库
type TaobaoRdsDbCreateAPIResponse struct {
	model.CommonResponse
	TaobaoRdsDbCreateAPIResponseModel
}

// TaobaoRdsDbCreateAPIResponseModel is rds创建数据库 成功返回结果
type TaobaoRdsDbCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"rds_db_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rds创建、查询、删除db返回结果的数据结构
	RdsDbInfo *RdsDbInfo `json:"rds_db_info,omitempty" xml:"rds_db_info,omitempty"`
}
