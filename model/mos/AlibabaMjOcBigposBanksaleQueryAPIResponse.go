package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjocbigposbanksalequeryAPIResponse 大pos银行卡查账接口 API返回值
// alibaba.mj.oc.bigpos.banksale.query
//
// 大pos银行卡查账接口，给收银员查询银行卡销售记录，便于调账
type AlibabamjocbigposbanksalequeryAPIResponse struct {
	model.CommonResponse
	AlibabamjocbigposbanksalequeryAPIResponseModel
}

// AlibabamjocbigposbanksalequeryAPIResponseModel is 大pos银行卡查账接口 成功返回结果
type AlibabamjocbigposbanksalequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_oc_bigpos_banksale_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据主体
	Datas []AlibabamjocbigposbanksalequeryData `json:"datas,omitempty" xml:"datas>alibabamjocbigposbanksalequery_data,omitempty"`
	// 明细数量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}
