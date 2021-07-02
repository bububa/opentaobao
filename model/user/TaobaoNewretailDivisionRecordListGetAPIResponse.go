package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoNewretailDivisionRecordListGetAPIResponse 导购分佣明细列表 API返回值
// taobao.newretail.division.record.list.get
//
// 提供分页查询导购分佣明细的能力
type TaobaoNewretailDivisionRecordListGetAPIResponse struct {
	model.CommonResponse
	TaobaoNewretailDivisionRecordListGetAPIResponseModel
}

// TaobaoNewretailDivisionRecordListGetAPIResponseModel is 导购分佣明细列表 成功返回结果
type TaobaoNewretailDivisionRecordListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"newretail_division_record_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否有上一页
	HasPrevPage bool `json:"has_prev_page,omitempty" xml:"has_prev_page,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 返回代码
	ResCode string `json:"res_code,omitempty" xml:"res_code,omitempty"`
	// 数据列表
	DataList []TaobaoNewretailDivisionRecordListGetT `json:"data_list,omitempty" xml:"data_list>taobao_newretail_division_record_list_get_t,omitempty"`
	// 是否有下一页
	HasNextPage bool `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 具体返回讯息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
