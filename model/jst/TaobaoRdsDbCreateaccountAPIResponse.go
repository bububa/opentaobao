package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaordsdbcreateaccountAPIResponse rds创建数据库账户 API返回值
// taobao.rds.db.createaccount
//
// rds创建数据库账户
type TaobaordsdbcreateaccountAPIResponse struct {
	model.CommonResponse
	TaobaordsdbcreateaccountAPIResponseModel
}

// TaobaordsdbcreateaccountAPIResponseModel is rds创建数据库账户 成功返回结果
type TaobaordsdbcreateaccountAPIResponseModel struct {
	XMLName xml.Name `xml:"rds_db_createaccount_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaordsdbcreateaccountResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
