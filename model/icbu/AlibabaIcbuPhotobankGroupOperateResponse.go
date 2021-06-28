package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图片银行分组操作接口 APIResponse
alibaba.icbu.photobank.group.operate

修改用户图片银行的分组信息，包括 新增分组，删除分组，分组重命名
*/
type AlibabaIcbuPhotobankGroupOperateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_icbu_photobank_group_operate_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回的数据结果
    
    PhotoGroupResult   *PhotoGroupResult `json:"photo_group_result,omitempty" xml:"