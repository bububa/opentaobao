package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
阿里金融用户授权 APIResponse
alipay.xiaodai.user.permit

阿里金融为用户开通消息通道接口
*/
type AlipayXiaodaiUserPermitAPIResponse struct {
    model.CommonResponse
    Response *AlipayXiaodaiUserPermitResponse `json:"alipay_xiaodai_user_permit_response,omitempty"`
}

type AlipayXiaodaiUserPermitResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
