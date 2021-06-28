package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序添加用户到指定的活动 APIResponse
taobao.jst.miniapp.crowd.user.add

小程序添加用户到指定的活动
*/
type TaobaoJstMiniappCrowdUserAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jst_miniapp_crowd_user_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 添加成功
    
    Result   bool `json:"result,omitempty" xml:"