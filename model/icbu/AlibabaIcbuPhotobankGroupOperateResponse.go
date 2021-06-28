package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
图片银行分组操作接口 APIResponse
alibaba.icbu.photobank.group.operate

修改用户图片银行的分组信息，包括 新增分组，删除分组，分组重命名
*/
type AlibabaIcbuPhotobankGroupOperateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuPhotobankGroupOperateResponse `json:"alibaba_icbu_photobank_group_operate_response,omitempty"` 
    AlibabaIcbuPhotobankGroupOperateResponse
}

/* model for simplify = false
type AlibabaIcbuPhotobankGroupOperateResponse struct {

    // 接口返回的数据结果
    
    PhotoGroupResult  *struct {
        PhotoGroupResult  *PhotoGroupResult `json:"photo_group_result,omitempty"`
    } `json:"photo_group_result,omitempty"`
    

}
*/

type AlibabaIcbuPhotobankGroupOperateResponse struct {

    // 接口返回的数据结果
    PhotoGroupResult   *PhotoGroupResult `json:"photo_group_result,omitempty"`

}
