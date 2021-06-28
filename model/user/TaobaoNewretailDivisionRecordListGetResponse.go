package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
导购分佣明细列表 APIResponse
taobao.newretail.division.record.list.get

提供分页查询导购分佣明细的能力
*/
type TaobaoNewretailDivisionRecordListGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"newretail_division_record_list_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否有上一页
    
    HasPrevPage   bool `json:"has_prev_page,omitempty" xml:"