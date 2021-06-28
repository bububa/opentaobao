package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销活动用户上传 APIResponse
tmall.ccf.crowd.activityuser.upload

搜集ISV的活动用户信息，将其沉淀为活动人群数据
*/
type TmallCcfCrowdActivityuserUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_ccf_crowd_activityuser_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Data   bool `json:"data,omitempty" xml:"