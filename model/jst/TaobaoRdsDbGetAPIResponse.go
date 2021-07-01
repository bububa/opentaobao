package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdsDbGetAPIResponse
查询rds下的数据库 API返回值
taobao.rds.db.get

查询rds实例下的数据库 */
type TaobaoRdsDbGetAPIResponse struct {
	model.CommonResponse
	TaobaoRdsDbGetAPIResponseModel
}

// TaobaoRdsDbGetAPIResponseModel is 查询rds下的数据库 成功返回结果
type TaobaoRdsDbGetAPIResponseModel struct {
	XMLName xml.Name `xml:"rds_db_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据库查询返回数据结构
	RdsDbInfos []RdsDbInfo `json:"rds_db_infos,omitempty" xml:"rds_db_infos>rds_db_info,omitempty"`
}
