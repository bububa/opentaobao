package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaordsdbgetdbAPIResponse rds获取RDS的DB API返回值
// taobao.rds.db.getdb
//
// rds获取RDS的DB
type TaobaordsdbgetdbAPIResponse struct {
	model.CommonResponse
	TaobaordsdbgetdbAPIResponseModel
}

// TaobaordsdbgetdbAPIResponseModel is rds获取RDS的DB 成功返回结果
type TaobaordsdbgetdbAPIResponseModel struct {
	XMLName xml.Name `xml:"rds_db_getdb_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaordsdbgetdbResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
