package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
导购分佣明细列表 APIResponse
taobao.newretail.division.record.list.get

提供分页查询导购分佣明细的能力
*/
type TaobaoNewretailDivisionRecordListGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoNewretailDivisionRecordListGetResponse `json:"newretail_division_record_list_get_response,omitempty"` 
    TaobaoNewretailDivisionRecordListGetResponse
}

/* model for simplify = false
type TaobaoNewretailDivisionRecordListGetResponse struct {

    // 是否有上一页
    
    HasPrevPage   bool `json:"has_prev_page,omitempty"`
    

    // 总数
    
    Total   int64 `json:"total,omitempty"`
    

    // 返回代码
    
    ResCode   string `json:"res_code,omitempty"`
    

    // 数据列表
    
    DataList  struct {
        T  []T `json:"t,omitempty"`
    } `json:"data_list,omitempty"`
    

    // 是否有下一页
    
    HasNextPage   bool `json:"has_next_page,omitempty"`
    

    // 页码
    
    PageNo   int64 `json:"page_no,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 每页数量
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 具体返回讯息
    
    Message   string `json:"message,omitempty"`
    

}
*/

type TaobaoNewretailDivisionRecordListGetResponse struct {

    // 是否有上一页
    HasPrevPage   bool `json:"has_prev_page,omitempty"`

    // 总数
    Total   int64 `json:"total,omitempty"`

    // 返回代码
    ResCode   string `json:"res_code,omitempty"`

    // 数据列表
    DataList   []T `json:"data_list,omitempty"`

    // 是否有下一页
    HasNextPage   bool `json:"has_next_page,omitempty"`

    // 页码
    PageNo   int64 `json:"page_no,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 每页数量
    PageSize   int64 `json:"page_size,omitempty"`

    // 具体返回讯息
    Message   string `json:"message,omitempty"`

}
