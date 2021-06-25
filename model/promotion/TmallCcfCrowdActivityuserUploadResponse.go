package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销活动用户上传 APIResponse
tmall.ccf.crowd.activityuser.upload

搜集ISV的活动用户信息，将其沉淀为活动人群数据
*/
type TmallCcfCrowdActivityuserUploadAPIResponse struct {
    model.CommonResponse
    Response *TmallCcfCrowdActivityuserUploadResponse `json:"tmall_ccf_crowd_activityuser_upload_response,omitempty"`
}

type TmallCcfCrowdActivityuserUploadResponse struct {

    // 返回结果
    Data   bool `json:"data,omitempty"`

    // 错误码
    ECode   string `json:"e_code,omitempty"`

    // 错误信息
    EMsg   string `json:"e_msg,omitempty"`

    // 是否失败
    Failed   bool `json:"failed,omitempty"`

    // 是否成功
    Suc   bool `json:"suc,omitempty"`

}
