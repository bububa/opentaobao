package dmp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询人群服务 APIResponse
taobao.dmp.crowds.get

查询人群服务
*/
type TaobaoDmpCrowdsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoDmpCrowdsGetResponse `json:"taobao_dmp_crowds_get_response,omitempty"`
}

type TaobaoDmpCrowdsGetResponse struct {

    // 1
    Results   []DmpCrowdDTO `json:"results,omitempty"`

}
